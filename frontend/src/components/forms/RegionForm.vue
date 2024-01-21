<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { classifyStore } from '@/store/classify';
import { profileStore } from '@/store/profile';

const ModalWin = defineAsyncComponent(() => import('@components/layouts/ModalWin.vue'));
const SelectDiv = defineAsyncComponent(() => import('@components/elements/SelectDiv.vue'));
const BtnGroupForm = defineAsyncComponent(() => import('@components/elements/BtnGroupForm.vue'));

const storeClassify = classifyStore();
const storeProfile = profileStore();

</script>

<template>
  <ModalWin :id="'modalRegion'" :title ="'Изменить регион'" :size="'modal-md'">
    <form @submit.prevent="storeProfile.dataProfile.updateItem" class="form form-check" role="form">
      <SelectDiv :name="'region_id'" :label="'Регион'" :select="storeClassify.classData.regions"
                 :model="storeProfile.dataProfile.form['region_id']"
                 @input-event="storeProfile.dataProfile.form['region_id'] = $event.target.value"/>
      <BtnGroupForm>
        <button class="btn btn-primary btn-md" data-bs-dismiss="modal" name="submit" type="submit">Принять</button>
        <button class="btn btn-primary btn-md" name="reset" type="reset">Очистить</button>
      </BtnGroupForm>
    </form>
  </ModalWin>
</template>
