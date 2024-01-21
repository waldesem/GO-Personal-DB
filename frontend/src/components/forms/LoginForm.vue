<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { loginStore } from '@/store/login';

const InputLabel = defineAsyncComponent(() => import('@components/elements/InputLabel.vue'));

const storeLogin = loginStore();

</script>

<template>
  <div class="border border-primary rounded p-5 ">
    <h5>{{ storeLogin.userData.action === 'login' ? 'Вход в систему' : 'Изменить пароль' }}</h5>
    <div class="py-3">
      <form @submit.prevent="storeLogin.userData.submitLogin" class="form form-check" role="form">
        <InputLabel :label="'Логин'" :name="'username'" :need="true"
                    :max="'16'" :min="'3'" :clsInput="'col-lg-8'"
                    :place="'Логин'" :pattern="'[a-zA-Z]+'"
                    @input-event="storeLogin.userData.form['username'] = $event.target.value"/>
        <div class="mb-3 row">
          <label class="col-form-label col-lg-2" for="password">Пароль: </label>
          <div class="col-lg-8">
            <div class="input-group">
              <input autocomplete="current-password" class="form-control" required 
                    id="password" name="password" minlength="8" maxlength="16" 
                    placeholder="Пароль" pattern="[0-9a-zA-Z]+"
                    :type="storeLogin.userData.hidden ? 'password' : 'text'" 
                    v-model="storeLogin.userData.form['password']" >
              <span class="input-group-text">
                <a role="button" @click="storeLogin.userData.hidden = !storeLogin.userData.hidden">
                  <i :class="storeLogin.userData.hidden ? 'bi bi-eye' : 'bi bi-eye-slash'"></i>
                </a>
              </span>
            </div>
            <div v-show="storeLogin.userData.action === 'login'" class="py-2">
              <a class="link-primary" href="#" @click="storeLogin.userData.action = 'password'">
                Изменить пароль
              </a>
            </div>
          </div>
        </div>
        <div v-if="storeLogin.userData.action === 'password'">
          <InputLabel :label="'Новый'" :name="'new_pswd'" :need="true"
                      :max="'16'" :min="'8'" :clsInput="'col-lg-6'"
                      :place="'Новый пароль'" :pattern="'[0-9a-zA-Z]+'"
                      :typeof="storeLogin.userData.hidden ? 'password' : 'text'" 
                      @input-event="storeLogin.userData.form['new_pswd'] = $event.target.value"/>
          <InputLabel :label="'Повтор'" :name="'conf_pswd'" :need="true" 
                      :max="'16'" :min="'8'" :clsInput="'col-lg-6'"
                      :place="'Повторите новый пароль'" :pattern="'[0-9a-zA-Z]+'" 
                      :typeof="storeLogin.userData.hidden ? 'password' : 'text'" 
                      @input-event="storeLogin.userData.form['conf_pswd'] = $event.target.value"/>
        </div>
        <div class="offset-lg-2 col-lg-10">
          <button class="btn btn-primary btn-md" name="submit" type="submit">
            {{ storeLogin.userData.action === 'login' ? 'Войти' : 'Изменить' }}
          </button>              
            &nbsp;
          <button v-show="storeLogin.userData.action === 'password'" 
                  class="btn btn-secondary btn-md" 
                  type="button" @click="storeLogin.userData.action = 'login'">
            Отменить
          </button>
        </div>
      </form>
    </div>
  </div>
</template>