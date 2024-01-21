<script setup lang="ts">

import { ref } from 'vue';
import { onBeforeMount, defineAsyncComponent } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';
import { authStore } from '@/store/token';
import { server } from '@/utilities/utils';

const ModalWin = defineAsyncComponent(() => import('@components/layouts/ModalWin.vue'));
const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));

const storeAuth = authStore();

const props = defineProps({
  path: {
    type: Array<string>,
    default: () => ['']
  }
});

onBeforeMount(() => {
  fileManager.value.path = props.path.slice(0, -1);
  fileManager.value.openFolder();  
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  fileManager.value.clearValue();
  next()
});

  const fileManager = ref({
    path: Array<string>(),
    folders: Array<string>(),
    files: Array<string>(),
    action: '',
    select: false,
    selected: Array<string>(),
    copied: Array<string>(),
    form: '',
    item: '',
    cols: 10,
    get rowsFolders () {
      return Math.ceil(this.folders.length / this.cols)
    },
    get rowsFiles () {
      return Math.ceil(this.files.length / this.cols)
    },

    getFoldersFiles: async function () {
    try {
      const response = await storeAuth.axiosInstance.get(`${server}/manager`);
      const { path, dirs, files } = response.data;
      this.assignValue(path, dirs, files);
      this.clearValue();
    } catch (error) {
      console.error(error);
    }
  },

  openFolder: async function (flag: string= 'open', item: string = '') {
    try {
      const response = await storeAuth.axiosInstance.post(`${server}/manager/${flag}`, {
        'path': this.path,
        'item': item
      });
      const { path, dirs, files } = response.data;

      this.assignValue(path, dirs, files);
    } catch (error) {
      console.error(error);
    }
  }, 

  openFile: async function (file: string) {
    try {
      const response = await storeAuth.axiosInstance.post(`${server}/manager/download`, {
        'path': this.path,
        'item': file
      }, { responseType: 'blob' });
      
      const fileName = file.split('/').pop();
      const fileExtension = fileName?.split('.').pop();
      
      this.item = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = this.item;
      link.download = `${fileName}.${fileExtension}`;
      link.dispatchEvent(new MouseEvent('click'));
    } catch (error) {
      console.error(error);
    }
  },

  updateItem: async function () {
    try {
      const response = this.action === 'create'
        ? await storeAuth.axiosInstance.post(`${server}/manager/${this.action}`, {
          'path': this.path,
          })
        : await storeAuth.axiosInstance.post(`${server}/manager/${this.action}`, {
          'path': this.path,
          'old': this.selected[0],
          'new': this.form
        });
      const { path, dirs, files } = response.data;
      this.assignValue(path, dirs, files);
      this.clearValue();
    } catch (error) {
      console.error(error);
    }
  },

  copyItem: async function () {
    if (this.selected.length) {
      try {
        const response = await storeAuth.axiosInstance.post(`${server}/manager/${this.action}`, {
          'path': this.path,
          'old': this.copied,
          'new': this.selected,
        });
        const { path, dirs, files } = response.data;
        this.assignValue(path, dirs, files);
        this.clearValue();
      } catch (error) {
        console.error(error);
      }
    }
  },

  deleteItem: async function () {
    if (this.selected.length) {
      if (confirm("Вы действительно хотите удалить?")) {
        try {
          const response = await storeAuth.axiosInstance.post(`${server}/manager/delete`, {
            'path': this.path,
            'items': this.selected
          });
          const { path, dirs, files } = response.data;
          this.assignValue(path, dirs, files);
          this.clearValue();
        } catch (error) {
          console.error(error);
        }
      }
    }
  },

  assignValue: function (path: [], dirs: [], files: []) {
    this.path = path;
    this.folders = dirs;
    this.files = files;
  },

  clearValue: function () {
    this.form = "";
    this.select = false;
    this.selected = [];
    this.copied = [];
  },
});

function fileType(file: string): string {
  const fileExtension = file.split('.').pop();
  
  switch (fileExtension) {
    case 'txt':
      return 'bi-file-text';
    case 'json':
      return 'bi-file-code';
    case 'html':
      return 'bi-file-code';
    case 'pdf':
      return 'bi-file-pdf';
    case 'docx':
      return 'bi-file-word';
    case 'doc':
      return 'bi-file-word-fill';
    case 'rtf':
      return 'bi-earmark-word';
    case 'xlsx':
      return 'bi-file-excel';
    case 'xlsm':
      return 'bi-file-excel-fill';
    case 'png':
      return 'bi-filetype-png';
    case 'jpg':
      return 'bi-filetype-jpg';
    case 'bmp':
      return 'bi-filetype-bmp';
    case 'zip':
      return 'bi-file-zip';
    case 'msg':
      return 'bi-envelope-paper';
    default:
      return 'bi-file';
  }
}

</script>

<template>
  <div class="container py-3">
    
    <HeaderDiv page-header="Файловый менеджер" />

    <div class="border border-primary p-5">

      <div class="row border border-primary p-3">

        <div class="col-1">
          <button type="button" class="btn btn-outline-primary" 
            @click="fileManager.getFoldersFiles" >
            <i class="bi bi-house" title="Домой"></i>
          </button>
        </div>
        
        <div class="col-1" :disabled="!fileManager.path.length">
          <button type="button" class="btn btn-outline-primary" title="Родитель"
                  @click="fileManager.path = fileManager.path.slice(0, -1); 
                          fileManager.openFolder('parent')"
                  :disabled="fileManager.select">
            <i class="bi bi-arrow-90deg-up"></i>
          </button>
        </div>
        
        <div class="col-1">
          <button type="button" class="btn btn-outline-primary" 
            @click="fileManager.action = 'create'; 
                    fileManager.updateItem()"
            :disabled="fileManager.select">
            <i class="bi bi-plus-square" title="Создать папку"></i>
          </button>
        </div>

        <div class="col-1">
          <button type="button" class="btn btn-outline-primary"
                  @click="fileManager.select = !fileManager.select;
                  fileManager.selected = []">
            <i class="bi bi-check-square" title="Выбрать"></i>
          </button>
        </div>

        <div class="col-1">
          <button type="button" class="btn btn-outline-primary"
                  @click="fileManager.action = 'copy'; 
                          fileManager.copied = fileManager.path;
                          fileManager.select = false"
                  :disabled="!fileManager.select || fileManager.selected.length === 0">
            <i class="bi bi-clipboard" title="Копировать"></i>
          </button>
        </div>

        <div class="col-1">
          <button type="button" class="btn btn-outline-primary"
                  @click="fileManager.action = 'сut'; 
                          fileManager.copied = fileManager.path;
                          fileManager.select = false"
                  :disabled="!fileManager.select || fileManager.selected.length === 0">
            <i class="bi bi-scissors" title="Вырезать"></i>
          </button>
        </div>
          
        <div class="col-1">
          <button type="button" class="btn btn-outline-primary"
                  :disabled="fileManager.select || fileManager.selected.length === 0"
                  @click="fileManager.copyItem">
            <i class="bi bi-clipboard-fill" title="Вставить"></i>
          </button>
        </div>
          
        <div class="col-1">
          <button type="button" class="btn btn-outline-primary" 
                  @click="fileManager.action = 'rename'"
                  :disabled="!fileManager.select || fileManager.selected.length !== 1"
                  data-bs-toggle="modal" data-bs-target="#modalFile">
            <i class="bi bi-pencil" title="Переменовать"></i>
          </button>
        </div>
          
        <div class="col-1">
          <button type="button" class="btn btn-outline-primary"
                  :disabled="!fileManager.select || fileManager.selected.length === 0"
                  @click="fileManager.deleteItem">
            <i class="bi bi-trash" title="Удалить"></i>
          </button>
        </div>
      </div>
    
      <div class="py-3">
        <nav aria-label="breadcrumb" :disabled="fileManager.select">
          <ol class="breadcrumb">
            <li class="breadcrumb-item active" aria-current="page" 
                v-for="item, idx in fileManager.path" :key="item">
              <a href="#" @click="fileManager.path = fileManager.path.slice(0, idx); 
                                  fileManager.openFolder('open', item)">
                {{ item }}
              </a>
            </li>
          </ol>
        </nav>
      </div>

      <div id="fileManager">
        <div v-for="folder, idx in fileManager.folders" :key="idx"  class="row">
          <div class="item-wrapper fs-6">
            <input class="form-check-input" type="checkbox" v-if="fileManager.select" 
                  :value="folder" v-model="fileManager.selected">
            &nbsp;
            <a type="button" class="btn btn-link btn-lg text-decoration-none" 
                    @click="fileManager.openFolder('open', folder)"
                    :disabled="fileManager.select">
              <i class="bi bi-folder"></i>
              {{ folder }}
            </a>
          </div>
        </div>
      </div>
      <div v-for="file, idx in fileManager.files" :key="idx"  class="row">
        <div class="item-wrapper fs-6">
          <input class="form-check-input" type="checkbox" v-if="fileManager.select" 
                :value="file" v-model="fileManager.selected">
          &nbsp;
          <a type="button" class="btn btn-link btn-lg text-decoration-none" 
                  @click="fileManager.openFolder(file)"
                  :disabled="fileManager.select">
            <i :class="fileType(file)"></i>
            {{ file }}
          </a>
        </div>
      </div>

    </div>
    
    <ModalWin :id="'modalFile'" :title ="'Переменовать'" :size="'modal-md'">
      <form @submit.prevent="fileManager.updateItem" class="form form-check" role="form">
        <div class="row">
          <div class="col">
            <input class="form-control" id="name" maxlength="250" name="name" type="text"
            v-model="fileManager.form">
          </div>
          <div class="col">
            <button class="btn btn-primary btn-md" data-bs-dismiss="modal" name="submit" type="submit">
              Принять
            </button>
          </div>
        </div>
      </form>
    </ModalWin>
  </div>
</template>

<style scoped>

#fileManager {
  height: 75vh;
  max-height: 75vh;
  overflow-y: auto;
}
.item-wrapper {
  display: flex;
  align-items: center;
}
</style>