<script setup lang="ts">

import { ref, defineAsyncComponent } from 'vue';
import { profileStore } from '@/store/profile';
import { classifyStore } from '@/store/classify';
import { clearForm } from '@utilities/utils'

const SwitchBtnForm = defineAsyncComponent(() => import('@components/elements/SwitchBtnForm.vue'));
const CollapseDiv = defineAsyncComponent(() => import('@components/elements/CollapseDiv.vue'));
const ResumeDiv = defineAsyncComponent(() => import('@components/tabs/divs/ResumeDiv.vue'));
const StaffDiv = defineAsyncComponent(() => import('@components/tabs/divs/StaffDiv.vue'));
const DocumentDiv = defineAsyncComponent(() => import('@components/tabs/divs/DocumentDiv.vue'));
const AddressDiv = defineAsyncComponent(() => import('@components/tabs/divs/AddressDiv.vue'));
const ContactDiv = defineAsyncComponent(() => import('@components/tabs/divs/ContactDiv.vue'));
const RelationDiv = defineAsyncComponent(() => import('@components/tabs/divs/RelationDiv.vue'));
const WorkplaceDiv = defineAsyncComponent(() => import('@components/tabs/divs/WorkplaceDiv.vue'));
const AffilationDiv = defineAsyncComponent(() => import('@components/tabs/divs/AffilationDiv.vue'));
const ResumeForm = defineAsyncComponent(() => import('@components/forms/ResumeForm.vue'));
const RegionForm = defineAsyncComponent(() => import('@components/forms/RegionForm.vue'));
const StaffForm = defineAsyncComponent(() => import('@components/forms/StaffForm.vue'));
const DocumentForm = defineAsyncComponent(() => import('@components/forms/DocumentForm.vue'));
const AddressForm = defineAsyncComponent(() => import('@components/forms/AddressForm.vue'));
const ContactForm = defineAsyncComponent(() => import('@components/forms/ContactForm.vue'));
const RelationForm = defineAsyncComponent(() => import('@components/forms/RelationForm.vue'));
const WorkplaceForm = defineAsyncComponent(() => import('@components/forms/WorkplaceForm.vue'));
const AffilationForm = defineAsyncComponent(() => import('@components/forms/AffilationForm.vue'));

const storeProfile = profileStore();
const storeClassify = classifyStore();

const hiddenSendBtn = ref(false);
const hiddenDelBtn = ref(false);

hiddenSendBtn.value = (storeProfile.dataProfile.resume['status'] 
                        !== storeClassify.classData.status['new'] 
                      && storeProfile.dataProfile.resume['status'] 
                        !== storeClassify.classData.status['update']
                      && storeProfile.dataProfile.resume['status'] 
                        !== storeClassify.classData.status['repeat']) 
                    || storeProfile.dataProfile.spinner

hiddenDelBtn.value = storeProfile.dataProfile.resume['status']
                       === storeClassify.classData.status['finish']
                    || storeProfile.dataProfile.spinner

function switchForm(item: string){
  storeProfile.dataProfile.flag === item 
    ? storeProfile.dataProfile.flag = '' 
    : storeProfile.dataProfile.flag = item;

  storeProfile.dataProfile.flag === item 
    ? storeProfile.dataProfile.action = 'create' 
    : storeProfile.dataProfile.action = ''; 

  clearForm(storeProfile.dataProfile.form)
};

</script>

<template>
  <div class="py-3">
  
    <template v-if="storeProfile.dataProfile.flag === 'resume' 
                 && storeProfile.dataProfile.action === 'update'" >
      <ResumeForm />
    </template>

    <template v-else>
      <RegionForm />
      <ResumeDiv v-if="storeProfile.dataProfile.resume"/>
      <p v-else >Данные отсутствуют</p>
    </template>
        
    <h6>Должности
      <SwitchBtnForm :item="'staff'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'staff'">
      <StaffForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.staffs">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.staffs" :key="idx"
            :id="'staff' + idx" :idx="idx" :label="'Должность #' + (idx + 1)">
          <StaffDiv :item="item"/>
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>

    <h6>Документы
      <SwitchBtnForm :item="'document'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'document'">
      <DocumentForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.docums">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.docums" :key="idx" 
            :id="'docum' + idx" :idx="idx" :label="'Документ #' + (idx + 1)">
          <DocumentDiv :item="item"/>    
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>    

    <h6>Адреса
      <SwitchBtnForm :item="'address'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'address'">
      <AddressForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.addrs">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.addrs" :key="idx" 
            :id="'addr' + idx" :idx="idx" :label="'Адрес #' + (idx + 1)">
          <AddressDiv :item="item"/>  
        </CollapseDiv>
      </div>  
      <p v-else >Данные отсутствуют</p>
    </template>

    <h6>Контакты
      <SwitchBtnForm :item="'contact'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'contact'">
      <ContactForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.conts">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.conts" :key="idx" 
            :id="'cont' + idx" :idx="idx" :label="'Контакт #' + (idx + 1)">
          <ContactDiv :item="item"/> 
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>

    <h6>Работа
      <SwitchBtnForm :item="'workplace'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'workplace'">
      <WorkplaceForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.works">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.works" :key="idx" 
            :id="'work' + idx" :idx="idx" :label="'Работа #' + (idx + 1)">
          <WorkplaceDiv :item="item"/>
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>

    <h6>Аффилированность
      <SwitchBtnForm :item="'affilation'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'affilation'">
      <AffilationForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.affilation">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.affilation" :key="idx" 
            :id="'affil' + idx" :idx="idx" :label="'Аффилированность #' + (idx + 1)">
          <AffilationDiv :item="item"/>
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>

    <h6>Связи
      <SwitchBtnForm :item="'relation'" :switchForm="switchForm" 
                     :flag="storeProfile.dataProfile.flag"/>
    </h6>
    <template v-if="storeProfile.dataProfile.flag === 'relation'">
      <RelationForm />
    </template>
    <template v-else>
      <div v-if="storeProfile.dataProfile.relate">
        <CollapseDiv v-for="item, idx in storeProfile.dataProfile.relate" :key="idx" 
            :id="'relate' + idx" :idx="idx" :label="'Связь #' + (idx + 1)">
          <RelationDiv :item="item"/>
        </CollapseDiv>
      </div>
      <p v-else >Данные отсутствуют</p>
    </template>

    <div class="d-print-none py-3">
      <div class='btn-group' role="group">
        <button class="btn btn-outline-primary" 
                :disabled="hiddenSendBtn"
                @click="storeProfile.dataProfile.getItem('resume', 'send', storeProfile.dataProfile.candId)" >
            {{ !storeProfile.dataProfile.spinner ? 'Отправить на проверку' : '' }}
          <span v-if="storeProfile.dataProfile.spinner" class="spinner-border spinner-border-sm"></span>
          <span v-if="storeProfile.dataProfile.spinner" role="status">Отправляется...</span>
        </button>
        <button type="button" class="btn btn-outline-danger" 
                :disabled="hiddenDelBtn" 
                @click="storeProfile.dataProfile.deleteItem('resume', 'delete', 
                                                storeProfile.dataProfile.resume['id'])">
          Удалить анкету
        </button>
      </div>
    </div>

  </div>
</template>
