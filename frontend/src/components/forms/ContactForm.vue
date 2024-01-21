<script setup lang="ts">

import { computed } from 'vue';
import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

const view = computed(() => {
  if (storeProfile.dataProfile.form['view'] === 'Телефон') {
    return 'tel';
  } else if (storeProfile.dataProfile.form['view'] === 'E-mail') {
    return 'email';
  } else {
    return 'text';
  }
});

const selected_item = {
  phone: 'Телефон',
  email: 'E-mail',
  other: 'Другое'
};

</script>

<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <SelectDiv :name="'view'" :label="'Выбрать'" :select="selected_item"
               :model="storeProfile.dataProfile.form['view']"
               @input-event="storeProfile.dataProfile.form['view'] = $event.target.value"/>
    <InputLabel :name="'contact'" :label="'Контакт'" :typeof="view" :need="true"
                :model="storeProfile.dataProfile.form['contact']"
                @input-event="storeProfile.dataProfile.form['contact'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" name="reset" type="reset">Очистить</button>
    </BtnGroupForm>
  </form>
</template>