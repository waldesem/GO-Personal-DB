<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const RowDivSlot = defineAsyncComponent(() => import('@components/elements/RowDivSlot.vue'));

const storeProfile = profileStore();

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});

</script>


<template>
  <RowDivSlot :slotTwo="true"  :print="true">
    <template v-slot:divTwo>
      <a href="#" title="Удалить"
        @click="storeProfile.dataProfile.deleteItem(props.item['id'].toString(), 'poligraf')">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a href="#" title="Изменить"
         @click="storeProfile.dataProfile.openForm('poligraf', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Тема'" :value="props.item['theme']"/>
  <RowDivSlot :label="'Результат'" :value="props.item['results']"/>
  <RowDivSlot :label="'Сотрудник'" :value="props.item['officer']"/>
  <RowDivSlot :label="'Дата'" :value="new Date(String(item['deadline'])).
                                          toLocaleDateString('ru-RU')"/>   
  <RowDivSlot :slotOne="true" :print="true">
    <form class="form" enctype="multipart/form-data" role="form" 
          @change="storeProfile.dataProfile.submitFile($event, 'poligraf', props.item['id'].toString())">
      <input class="form-control" id="file" type="file" ref="file" multiple>
    </form>
  </RowDivSlot>
</template>
