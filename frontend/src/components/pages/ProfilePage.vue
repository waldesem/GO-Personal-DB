<script setup lang="ts">

import { defineAsyncComponent, onBeforeMount, ref } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';
import { useRoute } from 'vue-router';
import { profileStore } from '@/store/profile';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));
const PhotoCard = defineAsyncComponent(() => import('@components/layouts/PhotoCard.vue'));
const AnketaTab = defineAsyncComponent(() => import('@components/tabs/AnketaTab.vue'));
const CheckTab = defineAsyncComponent(() => import('@components/tabs/CheckTab.vue'));
const PoligrafTab = defineAsyncComponent(() => import('@components/tabs/PoligrafTab.vue'));
const InvestigateTab = defineAsyncComponent(() => import('@components/tabs/InvestigateTab.vue'));
const InquiryTab = defineAsyncComponent(() => import('@components/tabs/InquiryTab.vue'));

const storeProfile = profileStore();

const route = useRoute();

storeProfile.dataProfile.candId = route.params.id.toString();

const tabsObject = {
  anketaTab: ['Анкета', AnketaTab],
  сheckTab: ['Проверки', CheckTab],
  poligrafTab: ['Полиграф', PoligrafTab],
  investigateTab: ['Расследования', InvestigateTab],
  inquiryTab: ['Запросы', InquiryTab]
};

const printPage = ref(false);

onBeforeMount(async () => {
  Promise.all([
    await storeProfile.dataProfile.getProfile(),
    await storeProfile.dataProfile.getImage()
  ])
});

onBeforeRouteLeave((_to: any, _from: any, next: () => void) => {
  storeProfile.dataProfile.cancelEdit();
  next();
});

</script>

<template>
  <div class="container py-3">
    
    <PhotoCard :url="storeProfile.dataProfile.url"
               :param="['image', storeProfile.dataProfile.resume.id]"
               :func="storeProfile.dataProfile.submitFile"/>

    <HeaderDiv :page-header="storeProfile.dataProfile.resume['fullname']" />

    <div v-if="!printPage" class="nav nav-tabs nav-justified" role="tablist">
      <button v-for="(value, key) in tabsObject" :key="key"
              class="nav-link" :class="{'active': key === 'anketaTab'}"
              data-bs-toggle="tab" :data-bs-target="`#${key}`"
              type="button" role="tab">
        {{value[0]}}
      </button>
    </div>

    <div v-if="!printPage" class="tab-content">
      <div v-for="(value, key) in tabsObject" :key="key" :id="key"
          class="tab-pane fade py-1" role="tabpanel" 
          :class="{' show active': key === 'anketaTab'}" >
        <component :is="value[1]"></component>
      </div>
    </div>

    <div v-if="printPage">
      <div v-for="(value, key) in tabsObject" :key="key" :id="key">
        <component :is="value[1]"></component>
      </div>
    </div>

    <a href="#" class="d-print-none" @click="printPage = !printPage">
      <i class="bi bi-printer fs-1" title="Версия для печати"></i>
    </a>

  </div>
</template>

<style scoped>
.bi-printer {
  position: fixed;
  top: 80px;
  right: 40px;
  z-index: 9999;
  border-radius: 50%;
  padding: 10px;
  cursor: pointer;
}
</style>
