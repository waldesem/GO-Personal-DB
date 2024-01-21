<script setup lang="ts">
 
import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();


</script>
 
<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <InputLabel :name="'relation'" :label="'Тип связи'" :need="true"
                :model="storeProfile.dataProfile.form['relation']" 
                @input-event="storeProfile.dataProfile.form['relation'] = $event.target.value"/>
    <InputLabel :name="'relation_id'" :label="'ID связи'" :need="true"
                :model="storeProfile.dataProfile.form['relation_id']" 
                @input-event="storeProfile.dataProfile.form['relation_id'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" name="reset" type="reset">Очистить</button>
    </BtnGroupForm>
  </form>
</template>
