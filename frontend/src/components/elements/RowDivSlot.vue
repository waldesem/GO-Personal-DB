<script setup lang="ts">

const props = defineProps({
  slotOne: {
    type: Boolean,
    default: false
  },
  slotTwo: {
    type: Boolean,
    default: false
  },
  label: {
    type: String,
    default: ''
  },
  value: {
    type: [String, Number],
    default: ''
  },
  print: {
    type: Boolean,
    default: false
  }
});

</script>

<template>
  <div class="row mb-3" :class="{'d-print-none': props.print}">
    <div class="col-md-3">{{props.label}}  
      <slot v-if="slotOne" name="divOne"></slot>
    </div>
    <div class="col-md-9">{{props.value}}
      <slot v-if="slotTwo" name="divTwo"></slot>
    </div>
  </div>
</template>