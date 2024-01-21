<script setup lang="ts">

import { defineAsyncComponent } from 'vue';
import { contactStore } from '@/store/contacts';

const InputSmall = defineAsyncComponent(() => import('@components/elements/InputSmall.vue'));

const storeContact = contactStore();

</script>


<template>
  <form @submit.prevent="storeContact.contactData.updateContact(
        $event, storeContact.contactData.action, storeContact.contactData.id
      )" 
        class="form form-check">
    <div class="row">
      <InputSmall :cls="'col-md-3'" :name="'company'" place="Организация" :need="true" 
                  :lst="'companies'" :selects="storeContact.contactData.companies"
                  :model="storeContact.contactData.form['company']"
                  @input-event="storeContact.contactData.form['company'] = $event.target.value"/>
      <InputSmall :cls="'col-md-3'" :name="'city'" place="Город" :need="true"
                  :lst="'cities'" :selects="storeContact.contactData.cities"
                  :model="storeContact.contactData.form['city']"
                  @input-event="storeContact.contactData.form['city'] = $event.target.value"/>
      <InputSmall :cls="'col-md-2'" :name="'fullname'" place="Имя" :need="true"
                  :model="storeContact.contactData.form['fullname']"
                  @input-event="storeContact.contactData.form['fullname'] = $event.target.value"/>
      <InputSmall :cls="'col-md-2'" :name="'phone'" place="Телефон" :pattern="'[0-9()]+'"
                  :model="storeContact.contactData.form['phone']"
                  @input-event="storeContact.contactData.form['phone'] = $event.target.value"/>
      <InputSmall :cls="'col-md-2'" :name="'adding'" place="Добав" :pattern="'[0-9]+'"
                  :model="storeContact.contactData.form['adding']"
                  @input-event="storeContact.contactData.form['adding'] = $event.target.value"/>    
    </div>
    <div class="row py-2">
      <InputSmall :cls="'col-md-3'" :name="'mobile'" place="Мобильный" :pattern="'[0-9]+'"
                  :model="storeContact.contactData.form['mobile']"
                  @input-event="storeContact.contactData.form['mobile'] = $event.target.value"/>
      <InputSmall :cls="'col-md-2'" :name="'mail'" place="Почта" 
                  :pattern="'[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$'"
                  :model="storeContact.contactData.form['mail']"
                  @input-event="storeContact.contactData.form['mail'] = $event.target.value"/>
      <InputSmall :cls="'col-md-3'" :name="'comment'" place="Комментарий"
                  :model="storeContact.contactData.form['comment']"
                  @input-event="storeContact.contactData.form['comment'] = $event.target.value"/>
      <div class="col-md-1">
        <button class="btn btn-outline-primary btn-sm" type="submit">Принять</button>
      </div>
      <div class="col-md-1">
        <button class="btn btn-outline-primary btn-sm" 
                @click="storeContact.contactData.action = '';
                        storeContact.contactData.id= ''">Отмена
        </button>
      </div>
    </div>
  </form>
</template>
              
              