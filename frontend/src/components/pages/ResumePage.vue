<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';
import { profileStore } from '@/store/profile';
import { clearForm } from '@utilities/utils'

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const ResumeForm = defineAsyncComponent(() => import('@components/forms/ResumeForm.vue'));

const storeProfile = profileStore();
storeProfile.dataProfile.action = 'create';

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  clearForm(storeProfile.dataProfile.form);
  next()
});

</script>

<template>
  <div class="container py-3">
    <HeaderDiv :page-header="'Создать анкету'" />
    <form class="form form-check" enctype="multipart/form-data" role="form" 
            @change="storeProfile.dataProfile.submitFile($event, 'anketa')">
        <div class="mb-3 row">
        <label class="col-form-label col-lg-2" for="file">Загрузить файл</label>
        <div class="col-lg-10">
            <input class="form-control" id="file" type="file" accept=".json" ref="file">
        </div>
        </div>
    </form>
    <ResumeForm />
  </div>
</template>