<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const InquiryForm = defineAsyncComponent(() => import('@components/forms/InquiryForm.vue'));
const CollapseDiv = defineAsyncComponent(() => import('@components/elements/CollapseDiv.vue'));
const InquiryDiv = defineAsyncComponent(() => import('@components/tabs/divs/InquiryDiv.vue'));

const storeProfile = profileStore();

</script>

<template>
  <div class="py-3">
    <InquiryForm v-if="(storeProfile.dataProfile.action === 'update'
                        || storeProfile.dataProfile.action === 'create') 
                      && storeProfile.dataProfile.flag === 'inquiry'" />

    <div v-else>
      <div v-if="storeProfile.dataProfile.needs.length">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.needs" :key="idx" 
            :id="'inquiry' + idx" :idx="idx" :label="'Запрос #' + (idx + 1)">
          <InquiryDiv :item="item" />
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
      <div class="d-print-none py-3">
        <a class="btn btn-outline-primary" type="button"
          @click="storeProfile.dataProfile.openForm('inquiry', 'create')">Добавить запись
        </a>
      </div>
    </div>
  </div>
</template>