<script setup lang="ts">

import { defineAsyncComponent, onBeforeMount } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';
import { adminStore } from '@store/admins';
import { debounce } from '@utilities/utils';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const UserForm = defineAsyncComponent(() => import('@components/forms/UserForm.vue'));

const storeAdmin = adminStore();

const searchUsers = debounce(() => {
  storeAdmin.dataUsers.getUsers() 
  }, 500
);

onBeforeMount(async () => {
  storeAdmin.dataUsers.getUsers()
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  storeAdmin.dataUsers.users = [];
  storeAdmin.dataUsers.search = '';
  next()
});

</script>


<template>
  <div class="container py-3">
    <HeaderDiv :page-header="'Список пользователей'" :cls="'text-secondary'" />
    <form @input.prevent="searchUsers" class="form form-check" role="form">
      <div class="row py-3">
        <input class="form-control" id="fullusername" name="fullusername" type="text" 
               v-model="storeAdmin.dataUsers.search">
      </div>
    </form>
    <div class="overflow py-2">
      <table class="table table-responsive align-middle">
        <thead>
          <tr height="50px">
            <th width="5%">#</th>
            <th>Имя пользователя</th>
            <th width="25%">Логин</th>
            <th width="20%">Создан</th>
            <th width="20%">Вход</th>
          </tr>
        </thead>
        <tbody>
          <tr >
            <td colspan="5">
              <table class="table table-hover table-responsive align-middle no-bottom-border">
                <tbody>
                  <tr height="50px" v-for="user in storeAdmin.dataUsers.users" :key="user.id">
                    <td width="5%">{{ user.id }}</td>
                    <td>{{ user.fullname }}</td>
                    <td width="25%">
                      <router-link :to="{ name: 'user', params: { id: user.id } }">
                        {{ user.username }}
                      </router-link>
                    </td>
                    <td width="20%">{{ new Date(user.pswd_create).toLocaleString('ru-RU') }}</td>
                    <td width="20%">{{ new Date(user.last_login).toLocaleString('ru-RU') }}</td>
                  </tr>
                </tbody>
              </table>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <button class="btn btn-outline-secondary" 
          data-bs-toggle="modal" data-bs-target="#modalUser"
          @click="storeAdmin.dataUsers.action = 'create'">
      Добавить пользователя 
    </button>

    <UserForm />
    
  </div>
</template>


<style scoped>
.overflow {
  max-height: 75vh;
  overflow-y: auto;
}
.no-bottom-border td {
  border-bottom: none;
}
</style>