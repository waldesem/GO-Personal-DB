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
    <TextLabel :name="'info'" :label="'Информация'"
               :model="storeProfile.dataProfile.form['info']"
               @input-event="storeProfile.dataProfile.form['info'] = $event.target.value"/>
    <InputLabel :name="'initiator'" :label="'Инициатор'"
                :model="storeProfile.dataProfile.form['initiator']"
                @input-event="storeProfile.dataProfile.form['initiator'] = $event.target.value"/>
    <InputLabel :name="'source'" :label="'Источник'"
                :model="storeProfile.dataProfile.form['source']"
                @input-event="storeProfile.dataProfile.form['source'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary" type="submit">Принять</button>
      <button class="btn btn-outline-primary" type="reset">Очистить</button>
      <button class="btn btn-outline-primary" type="button" 
              @click="storeProfile.dataProfile.cancelEdit">Отмена</button>
    </BtnGroupForm>
  </form>
</template>