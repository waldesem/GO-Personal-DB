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
      <a href="#" @click="storeProfile.dataProfile.deleteItem('contact', props.item['id'].toString())" title="Удалить">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a class="btn btn-link" title="Изменить"
        @click= "storeProfile.dataProfile.openForm('contact', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Вид'" :value="props.item['view']"/>
  <RowDivSlot :label="'Контакт'" :value="props.item['contact']"/>
</template>