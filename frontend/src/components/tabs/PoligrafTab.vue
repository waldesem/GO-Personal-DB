<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';

const PoligrafForm = defineAsyncComponent(() => import('@components/forms/PoligrafForm.vue'));
const CollapseDiv = defineAsyncComponent(() => import('@components/elements/CollapseDiv.vue'));
const PoligrafDiv = defineAsyncComponent(() => import('@components/tabs/divs/InvestigateDiv.vue'));

const storeProfile = profileStore();

</script>

  <template>
  <div class="py-3">
    <PoligrafForm v-if="(storeProfile.dataProfile.action === 'update' 
                  || storeProfile.dataProfile.action === 'create') 
                  && storeProfile.dataProfile.flag === 'poligraf'"/>
    <div v-else>
      <div v-if="storeProfile.dataProfile.pfo.length">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.pfo" :key="idx" 
            :id="'poligraf' + idx" :idx="idx" :label="'Полиграф #' + (idx + 1)">
          <PoligrafDiv :item="item" />
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
      <div class="d-print-none py-3">
        <a class="btn btn-outline-primary" type="button"
          @click="storeProfile.dataProfile.openForm('poligraf', 'create')">Добавить запись
        </a>
      </div>
    </div>
  </div>
</template>