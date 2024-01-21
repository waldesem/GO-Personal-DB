<script setup lang="ts">

import { computed, onBeforeMount, ref, defineAsyncComponent } from 'vue';
import { classifyStore } from '@/store/classify';
import { authStore } from '@/store/token';
import { server } from '@/utilities/utils';

const HeaderDiv = defineAsyncComponent(() => import('@components/layouts/HeaderDiv.vue'));

const storeAuth = authStore();
const storeClassify = classifyStore();

const todayDate = new Date();

const tableData = ref({
  header: '',
  tableSummed: <Array<Record<string, number>>>([{}]),
  stat: {
    region_id: 1,
    checks: [], 
    start: new Date(
      todayDate.getFullYear(), 
      todayDate.getMonth(), 1).toISOString().slice(0,10),
    end: todayDate.toISOString().slice(0,10)
  },
  
  submitData: async function (): Promise<void> {
    const response = await storeAuth.axiosInstance.post(`${server}/information`, {
      'start': this.stat.start, 'end': this.stat.end, 'region_id': this.stat.region_id 
    });
    const { candidates } = response.data;
    this.header = storeClassify.classData.regions[this.stat.region_id];
    this.stat.checks = candidates;
  }
});

onBeforeMount(async () => {
  await tableData.value.submitData()
});

computed(() => {
  tableData.value.header = storeClassify.classData.regions[tableData.value.stat.region_id];
});

</script>

<template>
  <div class="container py-3">
    <HeaderDiv :page-header="`Статистика по региону ${tableData.header} 
              за периоди c ${tableData.stat.start} по ${tableData.stat.end}`" />
    

    <div class="py-3">
      <table class="table table-hover table-responsive align-middle">
        <caption>Решения по кандидатам</caption>
        <thead><tr><th width="45%">Решение</th><th>Количество</th></tr></thead>
        <tbody>
          <tr height="50px" v-for="(value, index) in Object.keys(tableData.tableSummed)" 
              :key="index">
            <td >{{ value }}</td><td>{{Object.values(tableData.tableSummed)[index]}}</td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <div class="py-3">
      <form @submit.prevent="tableData.submitData" class="form form-check" role="form">
          <div class="mb-3 row required">
            <label class="col-form-label col-md-2" for="region">Регион</label>
            <div class="col-md-2">
              <select @change="tableData.submitData" 
                      class="form-select" id="region" name="region" 
                      v-model="tableData.stat.region_id">
                <option :value="tableData.stat.region_id" selected>
                  {{ storeClassify.classData.regions[tableData.stat.region_id] }}</option>
                <option v-for="name, value in storeClassify.classData.regions" 
                        :key="value" :value="value">{{name}}</option>                
              </select>
            </div>
            <label class="col-form-label col-md-1" for="start">Период:</label>
            <div class="col-md-2">
              <input class="form-control" id="start" name="start" required type="date" 
                  v-model="tableData.stat.start">
            </div>
            <div class="col-md-2">
              <input class="form-control" name="end" required type="date" 
                  v-model=tableData.stat.end>
            </div>
            <div class="col-md-2">
              <button class="btn btn-primary btn-md" name="submit" type="submit">
                  Принять</button>
            </div>
          </div>
      </form>
    </div>
  </div>
</template>
