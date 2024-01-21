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
      <a href="#" @click="storeProfile.dataProfile.deleteItem('staff', props.item['id'].toString())" title="Удалить">
        <i class="bi bi-trash"></i>
      </a>
      <a class="btn btn-link" title="Изменить"
        @click= "storeProfile.dataProfile.openForm('staff', 'update', props.item['id'].toString(), item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Должность'" :value="props.item['position']"/>
  <RowDivSlot :label="'Департамент'" :value="props.item['department']"/>
</template>