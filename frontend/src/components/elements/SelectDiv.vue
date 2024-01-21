<script setup lang="ts">

const props = defineProps({
  isneed: {
    type: Boolean,
    default: true
  },
  name: String,
  label: String,
  model: {
    type: [String, Number],
    default: ''
  },
  select: {
    type: Object as () => Record<string, any>,
    default: () => []
  },
  selectValue: String,
  selectName: String
});

</script>

<template>
  <div class="mb-3 row">
    <label class="col-form-label col-lg-2" :for="props.name">
      {{props.label}}
    </label>
    <div class="col-lg-10">
      <select class="form-select" 
              :required="props.isneed"
              :id="props.name"
              :name="props.name"
              :value="props.model"
              @input="$emit('input-event', $event)">
        <option v-for="value, index in props.select" :key="index"
                :value="value[props.selectValue as string]">{{value[props.selectName as string]}}</option>
      </select>
    </div>
  </div>
</template>