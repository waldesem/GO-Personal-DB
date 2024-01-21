import { defineStore } from 'pinia';
import { ref } from 'vue';
import { authStore } from '@/store/token';
import { alertStore } from '@store/alert';
import { clearForm, server } from '@utilities/utils';


export const contactStore = defineStore('contactStore', () => {

  const storeAuth = authStore();
  const storeAlert = alertStore();
  

  const contactData = ref({
    contacts: [],
    companies: [],
    cities: [],
    page: 1,
    prev: false,
    next: false,
    id: '',
    action: '',
    search: '',
    form: <Record<string, any>>({}),

    getContacts: async function (page: number): Promise<void> {
      try {
        const response = await storeAuth.axiosInstance.post(
          `${server}/connects/${page}`, {
            'search': this.search
          }
        );
        const [ datas, has_prev, has_next, companies, cities ] = response.data;
        Object.assign(this, {
          contacts: datas,
          companies: companies.companies,
          cities: cities.cities,
          prev: has_prev.has_prev,
          next: has_next.has_next
        });
      } catch (error) {
        console.error(error);
      }
    },
    
    updateContact: async function (event: Event, action: string, id: string): Promise<void> {
      event.preventDefault();
      try {
        const response = action === 'create'
          ? await storeAuth.axiosInstance.post(
            `${server}/connect/`, this.form.value
            )
          : await storeAuth.axiosInstance.patch(`${server}/connect/${id}`, this.form);
        const { item_id } = response.data;
  
        const alert = {
          'create': ['alert-success', `Создан контакт ID ${item_id}`],
          'edit': ['alert-info', `Контакт ID ${item_id} обновлен`]
        };
        storeAlert.alertMessage.setAlert(alert[action as keyof typeof alert][0], 
                                  alert[action as keyof typeof alert][1]);
        this.getContacts(this.page);
        this.action = '';
        clearForm(this.form);
      } catch (error) {
        console.log(error)
      }
    },
    deleteContact: async function (id: string): Promise<void> {
      if (confirm("Вы действительно хотите удалить контакт?")) {
        try {
          const response = await storeAuth.axiosInstance.delete(`${server}/connect/${id}`);
          console.log(response.status);
          storeAlert.alertMessage.setAlert('alert-success', `Контакт с ID ${id} удален`);
          this.getContacts(this.page);
        } catch (error) {
          console.log(error)
          storeAlert.alertMessage.setAlert('alert-danger', `Ошибка при удалении контакта с ID ${id}`);
        }
      }
    }
  });
  return { contactData };
});