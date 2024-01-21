<script setup lang="ts">

const props = defineProps({
  need: {
    type: Boolean,
    default: false
  },
  pattern: {
    type: String,
    default: '.*'
  },
  cls: String,
  place: String,
  name: String,
  model: String,
  lst: String,
  selects: Array
});

</script>

<template>
  <div :class="cls">
      <input class="form-control" max="255" type="text"
            :id="props.name" 
            :name="props.name"
            :required="props.need"
            :pattern="props.pattern"
            :placeholder="props.place"
            :list="props.lst"
            v-model="props.model">
    <datalist v-if="props.lst" :id="props.lst">
      <option v-for="item in selects" :value="item"></option>
    </datalist>
  </div>
</template>