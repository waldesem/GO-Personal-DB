<script setup lang="ts">

import { ref, defineAsyncComponent, onBeforeMount } from 'vue'
import { onBeforeRouteLeave } from 'vue-router';
import { classifyStore} from '@store/classify'
import { authStore } from '@/store/token';
import { alertStore } from '@store/alert';
import { debounce, server } from '@utilities/utils';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const PageSwitcher = defineAsyncComponent(() => import('@components/layouts/PageSwitcher.vue'));

const storeClassify = classifyStore();
const storeAuth = authStore();
const storeAlert = alertStore();

onBeforeMount(() => {
  if (storeClassify.classData.tables.length) {
    tableData.value.table = storeClassify.classData.tables[0];
    tableData.value.getItem(1);
  }
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  Object.assign(tableData.value.table, {
    table: '',
    tableItem: [],
    search: '',
    currentPage: 1,
    hasNext: false,
    hasPrev: false
  });
  next()
});

const tableData = ref({
  table: '',
  item: [],
  search: '',
  currentPage: 1,
  hasNext: false,
  hasPrev: false,

  getItem: async function(page: number): Promise<void> {
    this.currentPage = page;
    try {
      const response = await storeAuth.axiosInstance.post(
        `${server}/table/${this.table}/${page}`, {
          'search': this.search
        }
      );
      const [ datas, metadata ] = response.data;
      this.item = datas;
      this.hasNext = metadata.has_next;
      this.hasPrev = metadata.has_prev;
      
    } catch (error) {
      storeAlert.alertMessage.setAlert('alert-warning', error as string);
    }
  },

  deleteItem: async function (idItem: string): Promise<void>{
    if (confirm(`Вы действительно хотите удалить запись?`)) {
      try {
        const response = await storeAuth.axiosInstance.delete(
          `${server}/table/${this.table}/${idItem}`);
        console.log(response.status);
        storeAlert.alertMessage.setAlert('alert-warning', 
                            `Запись ${idItem} из ${this.table} удалена`);
        this.getItem(this.currentPage);
      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-warning', error as string);
      }
    };
  }
});

const searchItem = debounce(() => {
  tableData.value.getItem(1); 
  }, 500
);

</script>


<template>
  <div class="container py-3">
    <HeaderDiv :page-header="'Таблицы'" :cls="'text-secondary'" />
    <div class="row py-3">
      <div class="col-md-3">
        <form class="form form-check" role="form">
          <select v-if="storeClassify.classData.tables.length"
                  class="form-select" id="region" name="region" 
                  v-model="tableData.table" 
                  @change="tableData.getItem(1)">
            <option v-for="table, index in storeClassify.classData.tables" 
                          :key="index" :value="table">
              {{ table }}
            </option>
          </select>
        </form>
      </div>
      <div class="col-md-8">
        <form @input="searchItem" class="form form-check" role="form">
          <input class="form-control" id="name" name="name" placeholder="Поиск ID" type="text" 
                 v-model="tableData.search">
        </form>
      </div>
    </div>
    <div v-if="tableData.item.length" class="table-responsive py-3">
      <table class="table align-middle">
        <thead> 
          <tr>
            <th v-for="key, index in Object.keys(tableData.item[0])" 
                       :key="index">{{key}}</th>
            <th width="5%">action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row, index in tableData.item" :key="index">
            <td v-for="val, index in Object.values(row)" 
                      :key="index">{{ val }}</td>
            <td>
              <a href="#" @click="tableData.deleteItem(row['id'])" title="Удалить">
                <i class="bi bi-trash"></i>
              </a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <PageSwitcher :has_prev = "tableData.hasPrev"
                  :has_next = "tableData.hasNext"
                  :switchPrev = "tableData.currentPage -1"
                  :switchNext = "tableData.currentPage +1" 
                  :switchPage = "tableData.getItem"/>
  </div>
</template>


<style scoped>
table td
{
  border-bottom: none;
}
</style>