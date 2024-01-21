<script setup lang="ts">

import { inject } from 'vue';
import { loginStore } from '@/store/login';
import { messageStore } from '@/store/messages';
import { timeSince } from '@utilities/utils';

const storeLogin = loginStore();
const storeMessage = messageStore();

const pageIdentity = inject('pageIdentity') as string;

let isStarted = false;
if (!isStarted) {
  isStarted = true;
  setInterval(storeMessage.updateMessages, 1000000);
};

</script>


<template>
  <nav v-if="pageIdentity" class="navbar navbar-expand navbar-nav mr-auto navbar-dark d-print-none" 
       :class="`${pageIdentity ==='admins' ? 'bg-secondary': 'bg-primary'}`">
    <div class="container">
      <a class="navbar-brand">{{ pageIdentity ? pageIdentity.toUpperCase() : '' }}</a>
      <div class="navbar-nav mr-auto collapse navbar-collapse" id="navbarContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          
          <template v-if="pageIdentity.toLowerCase() === 'admins'">
            <li class="nav-item">
              <router-link :to="{ name: 'users', params: { group: 'admins' } }" class="nav-link active" href="#">
                Пользователи
              </router-link>
            </li>
            
            <li class="nav-item">
              <router-link :to="{ name: 'table', params: { group: 'admins' } }" 
                           class="nav-link active" href="#">
                Таблицы
              </router-link>
            </li>
          </template>

          <template v-if="pageIdentity.toLowerCase() === 'staffsec'">
            <li class="nav-item">
              <router-link :to="{ name: 'persons', params: { group: 'staffsec' }}" 
                            class="nav-link active">
                Кандидаты
              </router-link>
            </li>
            <li class="nav-item">
              <router-link :to="{ name: 'resume', params: { group: 'staffsec' } }" 
                            class="nav-link active">
                  Создать
              </router-link>
            </li>
            <li class="nav-item">
                <router-link :to="{ name: 'information', params: { group: 'staffsec' } }" 
                             class="nav-link active">
                  Информация
                </router-link>
            </li>
            
            <li class="nav-item">
              <router-link :to="{name: 'contacts', params: { group: 'staffsec' }}"    
                           class="nav-link active">
                  Контакты
              </router-link>
            </li>

            <li class="nav-item">
              <router-link :to="{ name: 'manager', params: { group: 'staffsec'} }" 
                           class="nav-link active" href="#">
                Файлы
              </router-link>
            </li>

            <li class="nav-item dropdown">
              <a class="nav-link active dropdown-toggle" role="button" data-bs-toggle="dropdown" href="#">
                Сообщения
                <span class="position-absolute translate-middle badge rounded-pill text-bg-success">
                  {{ storeMessage.messageData.messages.length }}
                </span>
              </a>
                <ul class="dropdown-menu" id="messages">
                  <h6 class="dropdown-header">Новые сообщения</h6>
                  <li v-for="message in storeMessage.messageData.messages" :key="message['id']">
                    <a class="dropdown-item">
                      <p>{{ timeSince(message['create']) }}</p>
                      <p>{{ message['title'] }}</p>
                    </a>
                  </li>
                  <div class="dropdown-divider"></div>
                  <li>
                    <router-link :to="{ name: 'messages', params: { group: 'staffsec' } }"
                                  class="dropdown-item" >
                      Открыть сообщения
                    </router-link>
                  </li>
                </ul>
            </li>
          </template>

        </ul>
        
        <li class="nav-item dropdown d-flex">
          <a href="#" class="nav-link active dropdown-toggle" role="button" 
              data-bs-toggle="dropdown" :title="storeLogin.userData.fullName ? storeLogin.userData.fullName : ''">
            {{ storeLogin.userData.fullName 
              ? storeLogin.userData.fullName.split(' ').map(item => item.charAt(0)).join('') 
              : '' }}
            <i class="bi bi-person-circle"></i>
          </a>
          <ul class="dropdown-menu">
            <li>
              <a class="dropdown-item" href="#" @click="storeLogin.userData.userLogout">Выход</a>
            </li>
          </ul>
        </li>
      </div>
    </div>
  </nav>
    
</template>


<style scoped>
  #messages {
    max-width: 640px;
    max-height: 640px;
    overflow-y: auto;
    overflow-x: auto;
  }
  #messages::after {
    display: none;
  }
</style>