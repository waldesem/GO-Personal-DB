<script setup lang="ts">

const props =defineProps({
  item: String,
  flag: String,
  switchForm: {
    type: Function,
    required: true
  }
})

</script>

<template>
  <a class="btn btn-link" @click="props.switchForm('contact')" 
     :title="props.flag === 'contact' ? 'Закрыть форму' : 'Добавить информацию'">
    <i :class="props.flag === 'contact' ? 'bi bi-dash-circle' : 'bi bi-plus-circle'"></i>
  </a>
</template>