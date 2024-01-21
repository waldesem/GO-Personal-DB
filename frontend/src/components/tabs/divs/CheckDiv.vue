<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const RowDivSlot = defineAsyncComponent(() => import('@components/elements/RowDivSlot.vue'));

const storeProfile = profileStore();

const props = defineProps({
  item: {
    type: Object,
    required: true
  },
  hiddenDelBtn: {
    type: Boolean,
    default: false
  },
  hiddeEditBtn: {
    type: Boolean,
    default: false
  }
});

</script>

<template>
  <RowDivSlot :slotTwo="true"  :print="true">
    <template v-slot:divTwo>
      <a :hidden="hiddenDelBtn" href="#" title="Удалить"
         @click="storeProfile.dataProfile.deleteItem(props.item['id'].toString(), 'check')">
        <i class="bi bi-trash"></i>
      </a>&nbsp; &nbsp; &nbsp;
      <a :hidden="hiddeEditBtn" href="#" title="Изменить"
         @click="storeProfile.dataProfile.openForm('check', 'update', 
                                        props.item['id'].toString(), props.item)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Проверка по местам работы<'" :value="props.item['workplace']"/>
  <RowDivSlot :label="'Бывший работник МТСБ'" :value="props.item['employee']"/>
  <RowDivSlot :label="'Проверка паспорта'" :value="props.item['document']"/>
  <RowDivSlot :label="'Проверка ИНН'" :value="props.item['inn']"/>
  <RowDivSlot :label="'Проверка ФССП'" :value="props.item['debt']"/>
  <RowDivSlot :label="'Проверка банкротства'" :value="props.item['bankruptcy']"/>
  <RowDivSlot :label="'Проверка БКИ'" :value="props.item['bki']"/>
  <RowDivSlot :label="'Проверка судебных дел'" :value="props.item['courts']"/>
  <RowDivSlot :label="'Проверка аффилированности'" :value="props.item['affiliation']"/>
  <RowDivSlot :label="'Проверка по списку террористов'" :value="props.item['terrorist']"/>
  <RowDivSlot :label="'Проверка нахождения в розыске'" :value="props.item['mvd']"/>
  <RowDivSlot :label="'Проверка в открытых источниках'" :value="props.item['internet']"/>
  <RowDivSlot :label="'Проверка Кронос'" :value="props.item['cronos']"/>
  <RowDivSlot :label="'Проверка Крос'" :value="props.item['cros']"/>
  <RowDivSlot :label="'Дополнительная информация'" :value="props.item['addition']"/>
  <RowDivSlot :label="'ПФО'" :value="props.item['pfo']"/>
  <RowDivSlot :label="'Комментарии'" :value="props.item['comments']"/>
  <RowDivSlot :label="'Результат проверки'" :value="props.item['conclusion']"/>
  <RowDivSlot :label="'Сотрудник'" :value="props.item['officer']"/>
  <RowDivSlot :label="'Дата'" :value="new Date(String(item['deadline'])).
                                          toLocaleDateString('ru-RU')"/>
  <RowDivSlot :slotOne="true"  :print="true">
    <form class="form" enctype="multipart/form-data" role="form" 
          @change="storeProfile.dataProfile.submitFile($event, 'check', item['id'].toString())">
      <input class="form-control" id="file" type="file" ref="file" multiple>
    </form>
  </RowDivSlot>
</template>