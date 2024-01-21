<script setup lang="ts">

import { ref, defineAsyncComponent } from 'vue';
import { classifyStore } from '@store/classify';
import { profileStore } from '@/store/profile';

const CheckForm = defineAsyncComponent(() => import('@components/forms/CheckForm.vue'));
const CollapseDiv = defineAsyncComponent(() => import('@components/elements/CollapseDiv.vue'));
const CheckDiv = defineAsyncComponent(() => import('@components/tabs/divs/CheckDiv.vue'));
const RobotDiv = defineAsyncComponent(() => import('@components/tabs/divs/RobotDiv.vue'));

const storeClassify = classifyStore();
const storeProfile = profileStore();

const hiddenAddBtn = ref(false);
const hiddenDelBtn = ref(false);
const hiddenEditBtn = ref(false);

hiddenDelBtn.value = storeProfile.dataProfile.resume['status'] 
                    === storeClassify.classData.status['finish'] 
                  || storeProfile.dataProfile.resume['status'] 
                    === storeClassify.classData.status['robot'];

hiddenEditBtn.value = storeProfile.dataProfile.resume['status'] 
                      !== storeClassify.classData.status['save'] 
                    && storeProfile.dataProfile.resume['status'] 
                      !== storeClassify.classData.status['cancel'] 
                    && storeProfile.dataProfile.resume['status'] 
                      !== storeClassify.classData.status['manual']

hiddenAddBtn.value = ![storeClassify.classData.status['new'], 
                      storeClassify.classData.status['update'], 
                      storeClassify.classData.status['save'], 
                      storeClassify.classData.status['repeat']].
                        includes(storeProfile.dataProfile.resume['status'])
</script>

<template>
  <div class="py-3">
    <CheckForm v-if="storeProfile.dataProfile.action === 'update' 
                      && storeProfile.dataProfile.flag === 'check'" />
    <div v-else>
      <div v-if="storeProfile.dataProfile.verification.length">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.verification" :key="idx" 
            :id="'check' + idx" :idx="idx" :label="'Проверка #' + (idx + 1)">
          <CheckDiv :item="item" :hiddenDelBtn="hiddenDelBtn" :hiddeEditBtn="hiddenEditBtn" />
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
      <div class="d-print-none py-3">
        <a class="btn btn-outline-primary" type="button"
          @click="storeProfile.dataProfile.openForm('check', 'create')">Добавить запись
        </a>
      </div>
      <div v-if="storeProfile.dataProfile.robot.length">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.robot" :key="idx" 
            :id="'check' + idx" :idx="idx" :label="'Робот #' + (idx + 1)">
          <RobotDiv :item="item" />
        </CollapseDiv>
      </div>
    </div>
  </div>
</template>