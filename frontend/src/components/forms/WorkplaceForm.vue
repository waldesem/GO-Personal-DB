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
    <InputLabel :name="'start_date'" :label="'Начало работы'" :typeof="'date'"
                :model="storeProfile.dataProfile.form['start_date']"
                @input-event="storeProfile.dataProfile.form['start_date'] = $event.target.value"/>
    <InputLabel :name="'end_date'" :label="'Окончание работы'" :typeof="'date'"
                :model="storeProfile.dataProfile.form['end_date']"
                @input-event="storeProfile.dataProfile.form['end_date'] = $event.target.value"/>
    <InputLabel :name="'workplace'" :label="'Место работы'" :need="true"
               :model="storeProfile.dataProfile.form['workplace']"
               @input-event="storeProfile.dataProfile.form['workplace'] = $event.target.value"/>
    <TextLabel :name="'address'" :label="'Адрес организации'"
               :model="storeProfile.dataProfile.form['address']"
               @input-event="storeProfile.dataProfile.form['address'] = $event.target.value"/>
    <TextLabel :name="'position'" :label="'Должность'"
               :model="storeProfile.dataProfile.form['position']"
               @input-event="storeProfile.dataProfile.form['position'] = $event.target.value"/>
    <TextLabel :name="'reason'" :label="'Причина увольнения'"
               :model="storeProfile.dataProfile.form['reason']"
               @input-event="storeProfile.dataProfile.form['reason'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" name="reset" type="reset">Очистить</button>
    </BtnGroupForm>
  </form>
</template>
