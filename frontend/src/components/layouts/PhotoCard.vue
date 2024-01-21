<script setup lang="ts">

import { ref } from 'vue';

const props = defineProps({
  url: String,
  param: {
    type: Array<String>,
    required: true
  },
  func: {
    type: Function,
    required: true
  }
});

const toggleForm = ref({
  showForm: false,
  handleMouse: function () {
    this.showForm = !this.showForm;
  }
});

</script>


<template>
  <div class="card" style="width: 16rem;">
    <div class="card-img-container" 
        @mouseover="toggleForm.handleMouse" 
        @mouseout="toggleForm.handleMouse">
      <img :src="props.url ? props.url : '/no-photo.png'" 
            style="width: 100%; height: auto;" class="card-img-top" alt="...">
      <form :class="{ 'form-visible': toggleForm.showForm }" 
          @change="props.func($event, ...props.param)" class="form">
        <input class="form-control form-control-sm" id="formImage" type="file">                  
      </form>
    </div>
  </div>
</template>


<style scoped>
.card-img-container {
  position: relative;
  display: inline-block;
}

.form {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  visibility: hidden;
}

.form-visible {
  visibility: visible;
}
</style>
