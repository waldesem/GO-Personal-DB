<script setup lang="ts">

import { defineAsyncComponent } from 'vue';

const RowDivSlot = defineAsyncComponent(() => import('@components/elements/RowDivSlot.vue'));

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});

</script>

<template>
  <RowDivSlot :label="'Проверка по кадровым данным<'" :value="props.item['employee']"/>
  <RowDivSlot :label="'Проверка ИНН'" :value="props.item['inn']"/>
  <RowDivSlot :label="'Проверка ФССП'" :value="props.item['debt']"/>
  <RowDivSlot :label="'Проверка банкротства'" :value="props.item['bankruptcy']"/>
  <RowDivSlot :label="'Проверка БКИ'" :value="props.item['bki']"/>
  <RowDivSlot :label="'Проверка судебных дел'" :value="props.item['courts']"/>
  <RowDivSlot :label="'Проверка по списку террористов'" :value="props.item['terrorist']"/>
  <RowDivSlot :label="'Проверка нахождения в розыске'" :value="props.item['mvd']"/>
  <RowDivSlot :label="'Дата'" :value="new Date(String(item['deadline'])).
                                          toLocaleDateString('ru-RU')"/>
</template>