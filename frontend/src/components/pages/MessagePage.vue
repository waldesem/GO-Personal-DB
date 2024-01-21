<script setup lang="ts">

import { defineAsyncComponent, onBeforeMount } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';
import { messageStore } from '@/store/messages';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const PageSwitcher = defineAsyncComponent(() => import('@components/layouts/PageSwitcher.vue'));

const storeMessage = messageStore();

onBeforeMount(async () => {
  await storeMessage.updateMessages('all');
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  storeMessage.messageData.messages = [];
  storeMessage.messageData.currentPage = 1;
  storeMessage.messageData.hasNext = false;
  storeMessage.messageData.hasPrev = false;
  next();
});

</script>


<template>
  <div class="container py-3">
    <HeaderDiv :page-header="'Сообщения'" />
    <div class="row justify-content-between">
      <div class="col">
        <p>
          <a href="#" class="link-info" @click="storeMessage.updateMessages('read')">
            Отметить все прочитанными
        </a>
      </p>
      </div>
      <div class="col text-end">
        <p>
          <a href="#" class="link-danger" @click="storeMessage.updateMessages('delete')">
            Удалить все сообщения
          </a>
        </p>
      </div>
    </div>
    <div class="py-2">
      <table class="table table-responsive align-middle">
        <thead> 
          <tr>
            <th width="5%">#</th>
            <th width="15%">Категория</th>
            <th width="15%">Тема</th>
            <th>Сообщение</th>
            <th width="15%">Статус</th>
            <th width="10%">Дата</th>
          </tr>
        </thead>
        <tbody v-if="storeMessage.messageData.messages.length">
          <tr v-for="message in storeMessage.messageData.messages" :key="message['id']" >
            <td width="5%">{{ `#${message['id']}` }}</td>
            <td width="15%">{{ message['category'] }}</td>
            <td width="15%">{{ message['title'] }}</td>
            <td>{{ message['report'] }}</td>
            <td width="15%">{{ message['status'] }}</td>
            <td width="10%">{{ message['created'] }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <PageSwitcher :has_prev = "storeMessage.messageData.hasPrev"
                  :has_next = "storeMessage.messageData.hasNext"
                  :switchPrev = "storeMessage.messageData.currentPage - 1"
                  :switchNext = "storeMessage.messageData.currentPage + 1"
                  :switchPage= "storeMessage.updateMessages" 
                  :option="'all'"/>
  </div>
</template>