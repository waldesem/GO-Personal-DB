<script setup lang="ts">

import { defineAsyncComponent } from 'vue';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));

</script>

<template>
  <div class="container py-3">
    <HeaderDiv :page-header="'404 - Page Not Found'" />
    <p class="text-secondary py-3">
      We're sorry, but the page you're looking for no longer exists or has been moved to another location
    </p>
    <p class="text-secondary py-3">
      Please check the URL or navigate back to the homepage
    </p>
  </div>
</template>
