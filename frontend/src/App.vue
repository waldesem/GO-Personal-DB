<script setup lang="ts">

import { onBeforeMount, watch, ref, provide } from 'vue';
import { defineAsyncComponent } from 'vue';
import { useRoute } from 'vue-router';
import { loginStore } from '@/store/login';

const NavBar = defineAsyncComponent(() => import('@components/layouts/NavBar.vue'));
const AlertMessage = defineAsyncComponent(() => import('@components/layouts/AlertMessage.vue'));
const MessagesDiv = defineAsyncComponent(() => import('@components/layouts/MessagesDiv.vue'));
const FooterDiv = defineAsyncComponent(() => import('@components/layouts/FooterDiv.vue'));

const storeLogin = loginStore();

const route = useRoute();

const pageIdentity = ref('')

provide('pageIdentity', pageIdentity)

watch(() => route.params.group,
  newValue => {
    pageIdentity.value = newValue as string
  }, { immediate: true });

onBeforeMount(() => {
  storeLogin.userData.getAuth();
});

</script>


<template>
  <NavBar />
  <AlertMessage />
  <router-view></router-view>
  <MessagesDiv />
  <FooterDiv />
</template>


<style>
html, body {
  scrollbar-gutter: stable;
}
</style>