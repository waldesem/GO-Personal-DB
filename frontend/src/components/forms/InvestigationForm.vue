<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const TextLabel = defineAsyncComponent(() => import('@components/elements/TextLabel.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

</script>

<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <InputLabel :name="'theme'" :label="'Тема проверки'" :need="true"
                :model="storeProfile.dataProfile.form['theme']"
                @input-event="storeProfile.dataProfile.form['theme'] = $event.target.value"/>
    <TextLabel :name="'info'" :label="'Информация'"
               :model="storeProfile.dataProfile.form['info']"
               @input-event="storeProfile.dataProfile.form['info'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary" type="submit">Принять</button>
      <button class="btn btn-outline-primary" type="reset">Очистить</button>
      <button class="btn btn-outline-primary" type="button" 
              @click="storeProfile.dataProfile.cancelEdit">Отмена</button>
    </BtnGroupForm>
  </form>
</template>