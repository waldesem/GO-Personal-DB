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
      <a href="#" @click="storeProfile.dataProfile.deleteItem('relation', props.item['id'].toString())" title="Удалить">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a class="btn btn-link" title="Изменить"
        @click= "storeProfile.dataProfile.openForm('relation', 'update', props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Тип связи'" :value="props.item['relation']"/>
  <RowDivSlot :label="'Связь'" :slotTwo="true">
    <router-link :to="{ name: 'profile', params: { group: 'staffsec', id: String(props.item['relation_id']) } }">
      ID #{{ props.item['relation_id'] }}
    </router-link>
  </RowDivSlot>
</template>