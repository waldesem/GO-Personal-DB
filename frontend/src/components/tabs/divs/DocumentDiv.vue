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
      <a href="#" @click="storeProfile.dataProfile.deleteItem('document', props.item['id'].toString())" title="Удалить">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a class="btn btn-link" title="Изменить"
        @click= "storeProfile.dataProfile.openForm('document', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Вид документа'" :value="props.item['view']"/>
  <RowDivSlot :label="'Серия'" :value="props.item['series']"/>
  <RowDivSlot :label="'Номер'" :value="props.item['number']"/>
  <RowDivSlot :label="'Кем выдан'" :value="props.item['agency']"/>
  <RowDivSlot :label="'Дата выдачи'" :value="new Date(String(props.item['issue'])).
                                              toLocaleDateString('ru-RU')"/>
</template>