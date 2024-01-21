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
    <InputLabel :name="'position'" :label="'Должность'" :need="true"
                :model="storeProfile.dataProfile.form['position']"
                @input-event="storeProfile.dataProfile.form['position'] = $event.target.value"/>
    <TextLabel :name="'department'" :label="'Подраздление'"
                :model="storeProfile.dataProfile.form['department']"
                @input-event="storeProfile.dataProfile.form['department'] = $event.target.value"/>

    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" name="reset" type="reset">Очистить</button>
    </BtnGroupForm>
  </form>
</template>
