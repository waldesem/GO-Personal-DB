<script setup lang="ts">

const props = defineProps({
  pageHeader: String,
  cls: {
    type: String,
    default: 'text-primary'
  }
});

</script>

<template>
  <div class="text-opacity-85 py-5" :class="cls">
    <h3>{{ props.pageHeader }}</h3>
  </div>
</template>