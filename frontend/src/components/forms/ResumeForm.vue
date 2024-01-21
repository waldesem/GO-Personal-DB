<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const TextLabel = defineAsyncComponent(() => import('@components/elements/TextLabel.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

const select_items = {
  candidate: "Кандидат",
  suspict: "Проверяемый"
};

</script>

<template>
  <div class="py-3">
    <form @submit.prevent="storeProfile.dataProfile.submitResume" class="form form-check" role="form">
      <SelectDiv :name="'category'" :label="'Категория'" :select="select_items"
                 @input-event="storeProfile.dataProfile.form['category'] = $event.target.value"
                 :model="storeProfile.dataProfile.form['category']" />
      <InputLabel :isneed="true" :name="'fullname'" :label="'Полное ФИО*'"
                  @input-event="storeProfile.dataProfile.form['fullname'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['fullname']"/>
      <TextLabel :name="'previous'" :label="'Изменение имени'"
                  @input-event="storeProfile.dataProfile.form['previous'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['previous']"/>
      <InputLabel :isneed="true" :name="'birthday'" :label="'Дата рождения*'" 
                  :typeof="'date'" 
                  @input-event="storeProfile.dataProfile.form['birthday'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['previous']"/>
      <TextLabel :name="'birthplace'" :label="'Место рождения'" 
                  @input-event="storeProfile.dataProfile.form['birthplace'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['birthplace']"/>
      <InputLabel :name="'country'" :label="'Гражданство'" :max="'255'"
                  @input-event="storeProfile.dataProfile.form['country'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['country']"/>
      <InputLabel :name="'ext_country'" :label="'Двойное гражданство'" :max="'255'"
                  @input-event="storeProfile.dataProfile.form['ext_country'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['ext_country']"/>
      <InputLabel :name="'snils'" :label="'СНИЛС'" :pattern="'[0-9]{11}'"
                  @input-event="storeProfile.dataProfile.form['snils'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['snils']"/>
      <InputLabel :name="'inn'" :label="'ИНН'" :max="'12'" :pattern="'[0-9]{12}'"
                  @input-event="storeProfile.dataProfile.form['inn'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['inn']"/>
      <TextLabel :name="'education'" :label="'Образование'" 
                  @input-event="storeProfile.dataProfile.form['education'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['education']"/>   
      <InputLabel :name="'marital'" :label="'Семейнное положение'" :max="'255'"
                  @input-event="storeProfile.dataProfile.form['marital'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['marital']"/>
      <TextLabel :name="'addition'" :label="'Дополнительно'" 
                  @input-event="storeProfile.dataProfile.form['addition'] = $event.target.value"
                  :model="storeProfile.dataProfile.form['addition']"/>

      <BtnGroupForm>
        <button class="btn btn-outline-primary" type="submit">Принять</button>
        <button class="btn btn-outline-primary" type="reset">Очистить</button>
        <router-link v-if="storeProfile.dataProfile.action !== 'update'" 
                class="btn btn-outline-primary" type="button"
                :to="{ name: 'persons', params: { group: 'staffsec' }}">
                Отмена
        </router-link>
        <button v-else class="btn btn-outline-primary" type="button" 
                @click="storeProfile.dataProfile.cancelEdit">Отмена</button>
      </BtnGroupForm>
    </form>
  </div>
</template>
