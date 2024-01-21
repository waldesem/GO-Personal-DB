<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

const selected_item = {
  passport: "Паспорт гражданина России",
  foreign: "Иностранный докумен",
  others: "Другое"
};
</script>

<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <SelectDiv :name="'view'" :label="'Выбрать'" :select="selected_item"
               :model="storeProfile.dataProfile.form['view']"
               @input-event="storeProfile.dataProfile.form['view'] = $event.target.value"/>
    <InputLabel :name="'series'" :label="'Серия документа'"
                :model="storeProfile.dataProfile.form['series']"
                @input-event="storeProfile.dataProfile.form['series'] = $event.target.value"/>
    <InputLabel :name="'number'" :label="'Номер документа'" :need="true"
                :model="storeProfile.dataProfile.form['number']"
                @input-event="storeProfile.dataProfile.form['number'] = $event.target.value"/>
    <InputLabel :name="'agency'" :label="'Орган выдавший'"
                :model="storeProfile.dataProfile.form['agency']"
                @input-event="storeProfile.dataProfile.form['agency'] = $event.target.value"/>
    <InputLabel :name="'issue'" :label="'Дата выдачи'" :typeof="'date'"
                :model="storeProfile.dataProfile.form['issue']"
                @input-event="storeProfile.dataProfile.form['issue'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" name="reset" type="reset">Очистить</button>
    </BtnGroupForm>>
  </form>
</template>
