<script setup lang="ts">
 
import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const TextLabel = defineAsyncComponent(() => import('@components/elements/TextLabel.vue'));
const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));
const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeProfile = profileStore();

const selected_item = {
  state: "Являлся государственным/муниципальным служащим",
  official: "Являлся государственным должностным лицом",
  relatives: "Связанные лица работают в государственных организациях",
  commercial: "Участвует в деятельности коммерческих организаций",
};

</script>

<template>
  <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
    <SelectDiv :name="'view'" :label="'Тип участия'" :select="selected_item"
               :model="storeProfile.dataProfile.form['view']"
               @input-event="storeProfile.dataProfile.form['view'] = $event.target.value"/>
    <TextLabel :name="'name'" :label="'Организация'"
               :model="storeProfile.dataProfile.form['name']"
               @input-event="storeProfile.dataProfile.form['name'] = $event.target.value"/>
    <InputLabel :name="'inn'" :label="'ИНН'" :need="true"
                :model="storeProfile.dataProfile.form['inn']"
                @input-event="storeProfile.dataProfile.form['inn'] = $event.target.value"/>
    <TextLabel :name="'position'" :label="'Должность'"
               :model="storeProfile.dataProfile.form['position']"
               @input-event="storeProfile.dataProfile.form['position'] = $event.target.value"/>
    <BtnGroupForm>
      <button class="btn btn-outline-primary btn-md" name="submit" type="submit">Принять</button>
      <button class="btn btn-outline-primary btn-md" type="reset">Очистить</button>
    </BtnGroupForm>
  </form>
</template>
