import { ref } from 'vue'
import { defineStore } from 'pinia';
import { authStore } from '@/store/token';
import { alertStore } from '@store/alert';
import { server, clearForm } from '@utilities/utils';
import router from '@/router/router';


export const adminStore = defineStore('adminStore', () => {

  const storeAlert = alertStore();
  const storeAuth = authStore();

  interface Group {
    id: string,
    group: string
  };

  interface Role {
    id: string,
    role: string
  };

  interface User {
    id: string,
    fullname: string,
    username: string,
    email: string,
    region_id: string,
    pswd_create: string,
    pswd_change: string,
    last_login: string,
    roles: Role[],
    groups: Group[],
    blocked: string,
    attempt: string
  };

  const dataUsers = ref({
    id: '',
    action: '',
    flag: '',
    role: '',
    group: '',
    search: '',
    users: <User[]>([]),
    profile: <User>({}),
    form: <Record<string, any>>({}),

    getUsers: async function(){
      try {
        const response = await storeAuth.axiosInstance.post(`${server}/users`, {'fullname': this.search});
        this.users = response.data;
      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-success', error as string)
      }
    },
    
    submitUser: async function(): Promise<void>{
      try {  
        const response = this.action === 'edit' 
          ? await storeAuth.axiosInstance.patch(`${server}/user`, this.form)
          : await storeAuth.axiosInstance.post(`${server}/user`, this.form);
        
        if (this.action === 'edit') {
          this.profile = response.data;
          storeAlert.alertMessage.setAlert('alert-success', 'Пользователь успешно изменен')
        } else {
          dataUsers.value.users = response.data;
          storeAlert.alertMessage.setAlert('alert-success', 'Пользователь успешно создан')
        };
  
      } catch (error) {
        console.error(error);
        storeAlert.alertMessage.setAlert('alert-danger', 'Ошибка сохранения данных');
      };
      clearForm(this.form);
    },

    userAction: async function(action: String): Promise<void>{
      try {
        const response = await storeAuth.axiosInstance.get(`${server}/user/${action}/${this.id}`);
        this.profile = response.data;
        if (action === 'drop'){
          storeAlert.alertMessage.setAlert('alert-success', 'Пароль сброшен');
        } else if (action === 'block') {
          storeAlert.alertMessage.setAlert('alert-success', 
            `Пользователь ${this.profile.blocked ? 'заблокирован' : 'разблокирован'}`);
        };
      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-success', error as string)
      }
    },

    userDelete: async function(): Promise<void>{
      if (confirm("Вы действительно хотите удалить пользователя?")){
        try {
          const response = await storeAuth.axiosInstance.delete(`${server}/user/${this.id}`);
          this.profile = response.data;
          storeAlert.alertMessage.setAlert('alert-success', 'Пользователь удалён');
          router.push({ name: 'users' });
      
        } catch (error) {
          storeAlert.alertMessage.setAlert('alert-danger', error as string)
        }
      }
    },

    updateGroupRole: async function(action: string, item: string, value: string): Promise<void> {
      if (value !== '') {
        try {
          const response = action === 'add' 
            ? await storeAuth.axiosInstance.get(`${server}/${item}/${value}/${this.id}`)
            : await storeAuth.axiosInstance.delete(`${server}/${item}/${value}/${this.id}`);
          this.profile = response.data;
          
          storeAlert.alertMessage.setAlert(
            'alert-success', 
            `${item === 'role' ? "Роль" : "Группа"} ${value} ${action === 'add' ? "добавлена" : "удалена"}`
            );
        } catch (error) {
          storeAlert.alertMessage.setAlert('alert-danger', error as string);
        };
      }
    }
  });
  return {
    dataUsers
  };
});
