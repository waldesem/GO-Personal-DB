<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';
import { classifyStore } from '@/store/classify';

const RowDivSlot = defineAsyncComponent(() => import('@components/elements/RowDivSlot.vue'));

const storeProfile = profileStore();
const storeClassify = classifyStore();

</script>

<template>
  <RowDivSlot :slotTwo="true" :print="true">
    <template v-slot:slotTwo>
      <a href="#" title="Изменить"
         @click="storeProfile.dataProfile.openForm('resume', 'update', storeProfile.dataProfile.resume['id'], storeProfile.dataProfile.resume)">
        <i class="bi bi-pencil-square"></i>
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Категория'" :value="storeProfile.dataProfile.resume['category']"/>
  <RowDivSlot :label="'Регион'" :slotTwo="true">
    <template v-slot:slotTwo>
      <a href="#" data-bs-toggle="modal" data-bs-target="#modalRegion"
        @click="storeProfile.dataProfile.openForm('resume', 'location', storeProfile.dataProfile.resume['id'], storeProfile.dataProfile.resume)">
          {{ storeClassify.classData.regions[storeProfile.dataProfile.resume['region_id']] }}
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Фамилия Имя Отчество'" :value="storeProfile.dataProfile.resume['fullname']"/>
  <RowDivSlot :label="'Изменение имени'" :value="storeProfile.dataProfile.resume['previous']"/>
  <RowDivSlot :label="'Дата рождения'" :value="storeProfile.dataProfile.resume['birthday']"/>
  <RowDivSlot :label="'Место рождения'" :value="storeProfile.dataProfile.resume['birthplace']"/>
  <RowDivSlot :label="'Гражданство'" :value="storeProfile.dataProfile.resume['country']"/>
  <RowDivSlot :label="'Второе гражданство'" :value="storeProfile.dataProfile.resume['ext_country']"/>
  <RowDivSlot :label="'СНИЛС'" :value="storeProfile.dataProfile.resume['snils']"/>
  <RowDivSlot :label="'ИНН'" :value="storeProfile.dataProfile.resume['inn']"/>
  <RowDivSlot :label="'Образование'" :value="storeProfile.dataProfile.resume['education']"/>
  <RowDivSlot :label="'Дополнительная информация'" :value="storeProfile.dataProfile.resume['addition']"/>
  <RowDivSlot :label="'Материалы'" :slotTwo="true" :print="true">
    <template v-slot:divTwo>
    <router-link :to="{name: 'manager', params: {group: 'staffsec'}, query: {path: storeProfile.dataProfile.resume['path'].split('/')}}">
      {{ storeProfile.dataProfile.resume['path'] }}
    </router-link>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Статус'" :slotTwo="true">
    <template v-slot:divTwo>
      <a href="#" @click="storeProfile.dataProfile.getItem('resume', 'status', storeProfile.dataProfile.resume['id'])">
        {{ storeProfile.dataProfile.resume['status'] }}
      </a>
    </template>
  </RowDivSlot>
  <RowDivSlot :label="'Создан'" :value="new Date(String(storeProfile.dataProfile.resume['created'])).
                                              toLocaleDateString('ru-RU')"/>
  <RowDivSlot :label="'Обновлен'" :value="new Date(String(storeProfile.dataProfile.resume['updated'])).
                                              toLocaleDateString('ru-RU')"/>
  <RowDivSlot :label="'Внешний ID'" :value="storeProfile.dataProfile.resume['request_id']"/>
</template>
