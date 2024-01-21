<script setup lang="ts">

import { onBeforeRouteLeave } from 'vue-router';
import { defineAsyncComponent } from 'vue';
import { loginStore } from '@store/login';
import { alertStore } from '@store/alert';
import { clearForm } from '@utilities/utils';

const LoginForm = defineAsyncComponent(() => import('@components/forms/LoginForm.vue'));
const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));

const storeAlert = alertStore();
const storeLogin = loginStore();

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  clearForm(storeLogin.userData.form);
  next()
});

storeAlert.alertMessage.attr = 'alert-info';
storeAlert.alertMessage.text = 'Авторизуйтесь для входа в систему';

</script>


<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="align-middle col col-6">
        <HeaderDiv :page-header="'StaffSec - кадровая безопасность'" />
        <LoginForm />
      </div>
    </div>
  </div>
</template>


<style scoped>
.container {
  height: 75vh;
}
</style>