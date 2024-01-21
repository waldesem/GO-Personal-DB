<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const TextLabel = defineAsyncComponent(() => import('@components/elements/TextLabel.vue'));
const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

const selected_item = {
  candidate: "Проверка кандидата",
  check: "Служебная проверка",
  investigation: "Служебное расследование"
}
</script>

<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <SelectDiv :name="'theme'" :label="'Тема проверки'" :select="selected_item"
               :model="storeProfile.dataProfile.form['theme']"
               @input-event="storeProfile.dataProfile.form['theme'] = $event.target.value"/>
    <TextLabel :name="'results'" :label="'Результат'"
               :model="storeProfile.dataProfile.form['results']"
               @input-event="storeProfile.dataProfile.form['results'] = $event.target.value"/>
    <BtnGroupForm>
          <button class="btn btn-outline-primary" type="submit">Принять</button>
          <button class="btn btn-outline-primary" type="reset">Очистить</button>
          <button class="btn btn-outline-primary" type="button" 
                  @click="storeProfile.dataProfile.cancelEdit">Отмена</button>
    </BtnGroupForm>
  </form>
</template>