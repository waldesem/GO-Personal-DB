import axios from 'axios';
import { ref } from 'vue'
import { defineStore } from 'pinia'
import { authStore } from '@/store/token';
import { alertStore } from '@store/alert';
import { classifyStore } from '@store/classify';
import { server } from '@utilities/utils';
import router from '@router/router';


export const loginStore = defineStore('loginStore', () => {

  const storeAlert = alertStore();
  const storeClasses = classifyStore();
  const storeAuth = authStore();
  

  const userData = ref({
    fullName: '',
    userName: '',
    userRoles: [],
    userGroups: [],
    region_id: '',
    action: 'login',
    hidden: true,
    form: <Record<string, any>>{},  
    
    getAuth: async function (): Promise<void> {
      try {
        const response = await storeAuth.axiosInstance.get(`${server}/login`);
        const { fullname, username, roles, groups, region_id } = response.data;
        this.assignUserData(fullname, username, roles, groups, region_id);
        
        this.hasRole('admin') 
          ? router.push({ name: 'users', params: {group: 'admins'}}) 
          : router.push({ name: 'persons', params: {group: this.userGroups[0]['group'] }});
        storeClasses.classData.getClasses();
        storeAlert.alertMessage.setAlert();

      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-warning', error as string);
        this.userLogout();
      }
    },

    submitLogin: async function (): Promise<void> {

      if (this.action === 'password') {
        if (this.form['password'] === this.form['new_pswd']) {
          storeAlert.alertMessage.setAlert('alert-warning', 'Старый и новый пароли совпадают');
          return
        };
        if (this.form['conf_pswd'] !== this.form['new_pswd']) {
          storeAlert.alertMessage.setAlert('alert-warning', 'Новый пароль и подтверждение не совпадают');
          return
        }
        delete this.form['conf_pswd'];
      };

      try {
        const response = this.action === 'password'
          ? await axios.patch(`${server}/login`, this.form)
          : await axios.post(`${server}/login`, this.form);
        const { message, access_token, refresh_token } = response.data;

        switch (message) {
          case 'Authenticated':
            if (this.action === 'password') {
              this.action = 'login';
              storeAlert.alertMessage.attr = 'alert-success';
              storeAlert.alertMessage.text = 'Войдите с новым паролем';
              delete this.form['new_pswd']
              delete this.form['conf_pswd']
            } else {     
              localStorage.setItem('refresh_token', refresh_token);
              localStorage.setItem('access_token', access_token);
              this.getAuth();
            };
            break;

          case 'Overdue':
            this.action = 'password';
            storeAlert.alertMessage.setAlert('alert-warning', 'Пароль просрочен. Измените пароль');
            break;

          case 'Denied':
            this.action = 'login';
            storeAlert.alertMessage.setAlert('alert-danger', 'Неверный логин или пароль');
            break;
        };
      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-warning', error as string);
        this.userLogout();
        //  clearForm(this.form)
      };
    },

    userLogout: async function (): Promise<void>{
      try {
        const response = await storeAuth.axiosInstance.delete(`${server}/login`);
        console.log(response.data);

      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-warning', error as string);
      };

      localStorage.removeItem('access_token');
      localStorage.removeItem('refresh_token');
      
      this.assignUserData();

      router.push({ name: 'login' });
    },

    hasRole: function (role: string): boolean {
      return this.userRoles.some((r: { role: any; }) => r.role === role);
    },

    hasGroup: function (group: string): boolean {
      return this.userGroups.some((g: { group: any; }) => g.group === group);
    },

    assignUserData(name = '', user = '', roles = [], groups = [], id = '') {
      Object.assign(this, {
        fullName: name,
        userName: user,
        userRoles: roles,
        userGroups: groups,
        region_id: id
      });
    }
  });
  return { 
    userData
  }
});