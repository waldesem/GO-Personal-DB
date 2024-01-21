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
        @click="storeProfile.dataProfile.deleteItem('inquiry', props.item['id'].toString(), )">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a href="#" title="Изменить"
        @click="storeProfile.dataProfile.openForm('inquiry', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Информация'" :value="props.item['info']"/>
  <RowDivSlot :label="'Иннициатор'" :value="props.item['initiator']"/>
  <RowDivSlot :label="'Источник'" :value="props.item['source']"/>
  <RowDivSlot :label="'Сотрудник'" :value="props.item['officer']"/>
  <RowDivSlot :label="'Дата запроса'" :value="props.item['deadline']"/>        
</template>