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
  <RowDivSlot :slotTwo="true" :print="true">
    <template v-slot:divTwo>
      <a href="#" @click="storeProfile.dataProfile.deleteItem('workplace', props.item['id'].toString())" title="Удалить">
        <i class="bi bi-trash"></i>
      </a>
      <a class="btn btn-link" title="Изменить"
        @click= "storeProfile.dataProfile.openForm('workplace', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Начало работы'" :value="props.item['start_date']"/>
  <RowDivSlot :label="'Окончание работы'" :value="props.item['end_date']"/>
  <RowDivSlot :label="'Организация'" :value="props.item['workplace']"/>
  <RowDivSlot :label="'КоАдреснтакт'" :value="props.item['address']"/>
  <RowDivSlot :label="'Должность'" :value="props.item['position']"/>
</template>