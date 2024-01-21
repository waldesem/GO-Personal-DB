<script setup lang="ts">

import { defineAsyncComponent, onBeforeMount } from 'vue';
import { onBeforeRouteLeave, useRoute } from 'vue-router';
import { classifyStore } from '@/store/classify';
import { adminStore } from '@store/admins';
import { clearForm } from '@utilities/utils';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const RowDivSlot = defineAsyncComponent(() => import('@components/elements/RowDivSlot.vue'));
const UserForm = defineAsyncComponent(() => import('@components/forms/UserForm.vue'));
//const PhotoCard = defineAsyncComponent(() => import('@components/layouts/PhotoCard.vue'));

const storeClassify = classifyStore();
const storeAdmin = adminStore();

const route = useRoute();

storeAdmin.dataUsers.id = route.params.id.toString();

onBeforeMount(async () => {
  storeAdmin.dataUsers.userAction('view');
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  clearForm(storeAdmin.dataUsers.form);
  Object.assign(storeAdmin.dataUsers, {
    id: '',
    action: '',
    search: '',
    role: '',
    group: ''
  });
  next()
}); 

</script>


<template>
  <div class="container py-3">
    <HeaderDiv :page-header="storeAdmin.dataUsers.profile.fullname" :cls="'text-secondary'"/>
    <div class="py-3">
      <!--PhotoCard :profileId="storeAdmin.profileData.id" :imageUrl="storeAdmin.profileData.image"/-->
      <div>
        <RowDivSlot :label="'ID'" 
                    :value="storeAdmin.dataUsers.profile.id"/>
        <RowDivSlot :label="'Имя пользователя'" 
                    :value="storeAdmin.dataUsers.profile.fullname"/>
        <RowDivSlot :label="'Логин'" 
                    :value="storeAdmin.dataUsers.profile.username"/>
        <RowDivSlot :label="'E-mail'" 
                    :value="storeAdmin.dataUsers.profile.email"/>
        <RowDivSlot :label="'Создан'" 
                    :value="new Date(storeAdmin.dataUsers.profile.pswd_create).toLocaleString('ru-RU')"/>
        <RowDivSlot :label="'Изменен'" 
                    :value="new Date(storeAdmin.dataUsers.profile.pswd_change).toLocaleString('ru-RU')"/>
        <RowDivSlot :label="'Вход'" 
                    :value="new Date(storeAdmin.dataUsers.profile.last_login).toLocaleString('ru-RU')"/>
        <RowDivSlot :label="'Попыток входа'" 
                    :value="storeAdmin.dataUsers.profile.attempt"/>
        <RowDivSlot :label="'Блокировка'" 
                    :value="storeAdmin.dataUsers.profile.blocked ? 'Заблокирован' : 'Разблокирован'"/>
        <RowDivSlot :label="'Группы'" :slotTwo="true">
          <template v-slot:divTwo>
            <ul v-for="(group, index) in storeAdmin.dataUsers.profile.groups" :key=index>
              <li>{{ group['group'] }}
                <a href="#" @click="storeAdmin.dataUsers.updateGroupRole('delete', 'group', group['id'])">
                  <i class="bi bi-dash-circle"></i>
                </a>
              </li>
            </ul>
            <form class="form form-check" role="form">
              <select class="form-select" id="group" name="group" style="width: 30%;"
                      v-model="storeAdmin.dataUsers.group" 
                      @change="storeAdmin.dataUsers.updateGroupRole('add', 'group', storeAdmin.dataUsers.group)">
                <option value="" selected>
                  Добавить группу
                </option>
                <option v-for="(group, index) in storeClassify.classData.groups" :key="index" :value="group['id']">
                  {{ group['group'] }}
                </option>
              </select>
            </form>
          </template>
        </RowDivSlot>
        <RowDivSlot :label="'Роли'" :slotTwo="true">
          <template v-slot:divTwo>
            <ul v-for="(role, index) in storeAdmin.dataUsers.profile.roles" :key=index>
              <li>{{ role['role'] }}
                <a href="#" @click="storeAdmin.dataUsers.updateGroupRole('delete', 'role', role['id'])">
                  <i class="bi bi-dash-circle"></i>
                </a>
              </li>
            </ul>
            <form class="form form-check" role="form">
              <select class="form-select" id="role" name="role" style="width: 30%;"
                  v-model="storeAdmin.dataUsers.role" 
                  @change="storeAdmin.dataUsers.updateGroupRole('add', 'role', storeAdmin.dataUsers.role)">
                <option value="" selected>
                  Добавить роль
                </option>
                <option v-for="(role, index) in storeClassify.classData.roles" :key="index" :value="role['id']">
                  {{ role['role'] }}</option>
              </select>
            </form>
          </template>
        </RowDivSlot>
      </div>
      <div class="btn-group py-3" role="group">
        <button class="btn btn-outline-secondary" type="button" 
                data-bs-toggle="modal" data-bs-target="#modalUser"
                @click="storeAdmin.dataUsers.action = 'edit';
                storeAdmin.dataUsers.form = storeAdmin.dataUsers.profile">
          Изменить пользователя
        </button>
        <button @click="storeAdmin.dataUsers.userAction('block')" class="btn btn-outline-secondary">
          {{storeAdmin.dataUsers.profile.blocked ? "Разблокировать" : 'Заблокировать' }}
        </button>
        <button @click="storeAdmin.dataUsers.userAction('drop')" type="button" class="btn btn-outline-secondary">
          Сбросить пароль
        </button>
        <button @click="storeAdmin.dataUsers.userDelete" type="button" class="btn btn-outline-secondary">
          Удалить
        </button>
      </div>
    </div>
    <UserForm />
  </div>
</template>


<style scoped>
ul, li {
  padding-left: 0;
}
</style>