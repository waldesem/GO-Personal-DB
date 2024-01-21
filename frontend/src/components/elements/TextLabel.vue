<script setup lang="ts">

const props = defineProps({
  name: String,
  label: String,
  model: {
    type: [String, Number],
    default: ''
  },
});

</script>

<template>
  <div class="mb-3 row">
    <label class="col-form-label col-lg-2" :for="props.name">
      {{props.label}}
    </label>
    <div class="col-lg-10">
      <textarea class="form-control" 
                :id="props.name" 
                :name="props.name"
                :placeholder="props.label"
                :value="props.model"
                @input="$emit('input-event', $event)">
      </textarea>
    </div>
  </div>
</template>