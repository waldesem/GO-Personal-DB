<script setup lang="ts">

import { ref } from 'vue';

const props = defineProps({
  id: String,
  idx: {
    type: Number,
    required: true
  },
  label: String
});

const toggle = ref(true);

</script>


<template>
  <div class="mb-3 py-2">
    <a class="link-primary" role="button" data-bs-toggle="collapse" :href="`#${props.id}`"
      @click="toggle = !toggle">
      <i class="bi" :class="toggle ? 'bi-plus-circle' : 'bi-dash-circle'"></i>
        {{ props.label }}
    </a>
    <div class="collapse" :class="{'show': props.idx == 0}" :id="props.id">
      <slot></slot>
    </div>
  </div>
</template>
