import { ref } from 'vue';
import { defineStore } from 'pinia';
import { classifyStore } from '@store/classify';
import { authStore } from '@/store/token';
import { alertStore } from '@store/alert';
import { server, clearForm } from '@utilities/utils';
import router from '@/router/router';

export const profileStore = defineStore('profileStore', () => {
  
  const storeAuth = authStore();
  const storeAlert = alertStore();
  const classifyApp = classifyStore();
  
  interface Resume {
    id: string;
    category: string;
    region_id: string;
    fullname: string;
    previous: string;
    birthday: string;
    birthplace: string;
    country: string;
    ext_country: string;
    snils: string;
    inn: string;
    education: string;
    marital: string;
    addition: string;
    path: string;
    status: string;
    created: string;
    updated: string;
    request_id: string;
  };
  
  interface Document {
    id: string;
    view: string;
    series: string;
    number: string;
    agency: string;
    issue: string;
  };
  
  interface Address {
    id: string;
    view: string;
    region: string;
    address: string;
  };
  
  interface Contact {
    id: string;
    view: string;
    contact: string;
  };
  
  interface Work {
    id: string;
    start_date: string;
    end_date: string;
    workplace: string;
    address: string;
    reason: string;
    position: string;
  };
  
  interface Staff {
    id: string;
    position: string;
    department: string;
  };
  
  interface Relation {
    id: string;
    relation: string;
    relation_id: string;
  };
  
  interface Affilation {
    id: string,
    view: string,
    name: string,
    inn: string,
    position: string,
    deadline: string
  };
  
  interface Verification {
    id: string;
    workplace: string;
    employee: string;
    document: string;
    inn: string;
    debt: string;
    bankruptcy: string;
    bki: string;
    courts: string;
    affiliation: string;
    terrorist: string;
    mvd: string;
    internet: string;
    cronos: string;
    cros: string;
    addition: string;
    pfo: string,
    conclusion: string;
    comments: string;
    deadline: string;
    officer: string;
  };

  interface Robot {
    id: string;
    employee: string;
    document: string;
    inn: string;
    debt: string;
    bankruptcy: string;
    bki: string;
    courts: string;
    affiliation: string;
    terrorist: string;
    mvd: string;
    deadline: string;
  }
  
  interface Pfo {
    id: string;
    theme: string;
    results: string;
    officer: string;
    deadline: string;
  };
  
  interface Inquisition {
    id: string;
    theme: string;
    info: string;
    officer: string;
    deadline: string;
  };
  
  interface Needs {
    id: string;
    info: string;
    initiator: string;
    source: string;
    officer: string;
    deadline: string;
  };
  
  const dataProfile = ref({
    candId: '',
    itemId: '',
    flag: '',
    action: '',
    url: '',
    spinner: false,
    resume: <Resume>({}),
    docums: Array<Document>(),
    addrs: Array<Address>(),
    conts: Array<Contact>(),
    works: Array<Work>(),
    staffs: Array<Staff>(),
    relate: Array<Relation>(),
    affilation: Array<Affilation>(),
    verification: Array<Verification>(),
    robot: Array<Robot>(),
    pfo: Array<Pfo>(),
    inquisition: Array<Inquisition>(),
    needs: Array<Needs>(),
    form: <Record<string, any>>({}),

    // getProfile: async function () {
    //   await Promise.all([
    //     [
    //       'resume', 
    //       'staff', 
    //       'document', 
    //       'address',
    //       'contact', 
    //       'workplace', 
    //       'relation',
    //       'affilation', 
    //       'check', 
    //       'robot',
    //       'poligraf', 
    //       'investigation', 
    //       'inquiry',
    //     ].map(async (item) => await this.getItem(item, 'view', this.candId))
    //   ]);
    // },
    
    getProfile: async function () {
      const response = await storeAuth.axiosInstance.get(`${server}/person/${this.candId}`);
      const { resume, staffs, documents, addresses, contacts, workplaces, 
        relations, affilations, checks, robots, poligrafs, investigations, inquiries }
        = response.data;
        this.resume = resume;
        this.staffs = staffs;
        this.docums = documents;
        this.addrs = addresses;
        this.conts = contacts;
        this.works = workplaces;
        this.relate = relations;
        this.affilation = affilations;
        this.verification = checks;
        this.robot = robots;
        this.pfo = poligrafs;
        this.inquisition = investigations;
        this.needs = inquiries;
    },

    getItem: async function (
      item: string, action: string, id: string): Promise<void> {
  
      if (item === 'check' && action === 'add'){
        if (this.resume['status'] === classifyApp.classData.status['save'] || 
            this.resume['status'] === classifyApp.classData.status['manual'] ||
            this.resume['status'] === classifyApp.classData.status['robot']) {
          
          storeAlert.alertMessage.setAlert('alert-warning', 
                                            'Нельзя добавить проверку к текущему статусу');
          return
        };
      };
      if (item === 'check' && action === 'self'){
        if (!confirm('Вы действительно делегировать анкету себе?')) {
        return
        };
      };
      if (item === 'resume' && action === 'status'){
        if (!confirm('Вы действительно хотите изменить статус резюме]?')) {
        return
        };
      };
      try {
        const response = await storeAuth.axiosInstance.get(`${server}/${item}/${action}/${id}`);
        switch (item){
          case 'resume':
            this.resume = response.data;
            break;
          case 'staff': 
            this.staffs = response.data;
            break;
          case 'document': 
            this.docums = response.data;
            break;
          case 'address': 
            this.addrs = response.data;
            break;
          case 'contact': 
            this.conts = response.data;
            break;
          case 'workplace': 
            this.works = response.data;
            break;
          case 'relation': 
            this.relate = response.data;
            break;
          case 'affilation': 
            this.affilation = response.data;
            break;
          case 'check': 
            this.verification = response.data;
            break;
          case 'robot':
            this.robot = response.data;
            break;
          case 'poligraf': 
            this.pfo = response.data;
            break;
          case 'investigation':
            this.inquisition = response.data;
             break;
          case 'inquiry': 
            this.needs = response.data;
            break;
          default:
             console.log(response.data);
            break;
        };
        if (action === 'status'){
          storeAlert.alertMessage.setAlert('alert-info', 'Статус анкеты обновлен');
        
        } else if (action === 'send'){
          storeAlert.alertMessage.setAlert('alert-success', 'Анкета отправлена на проверку');
          this.spinner = false
          window.scrollTo(0, 0);
          this.getItem('check', 'get', this.candId);
        
        } else if (item === 'check' && (action === 'add' || action === 'self')){
          this.getItem('check', 'get', this.candId);
        }
  
      } catch (error) {
        console.error(error)
        storeAlert.alertMessage.setAlert('alert-danger', `Ошибка обработки ${error}`);
      }
    },
  
    updateItem: async function(): Promise<void> {
      this.flag === 'registry' ? this.spinner = true : this.spinner = false;
        try {
        const response = this.action === 'create' 
          ? await storeAuth.axiosInstance.post(`${server}/${this.flag}/${this.action}/${this.candId}`, this.form)
          : await storeAuth.axiosInstance.patch(`${server}/${this.flag}/${this.action}/${this.itemId}`, this.form);
  
        console.log(response.status);
  
        storeAlert.alertMessage.setAlert('alert-success', 'Данные успешно обновлены');
        
        if (['check', 'poligraf'].includes(this.flag)) {
          this.getItem('resume', 'get', this.candId)
        };
        this.getItem(this.flag, this.action, this.candId);
  
      } catch (error) {
        storeAlert.alertMessage.setAlert('alert-danger', `Возникла ошибка ${error}`);
      }
      clearForm(this.form);
      this.action = '';
      this.flag = '';
      this.spinner = false;
    },
    
    deleteItem: async function (
      item: string, id: string, action: string = 'delete'): Promise<void> {
      if ([classifyApp.classData.status['robot'], 
           classifyApp.classData.status['finish']].includes(this.resume['status']) 
          && (item === 'check' || item === 'resume')) {
        storeAlert.alertMessage.setAlert('alert-warning', 
                                         'Нельзя удалить запись с текущим статусом');
        return
      };
      if (confirm(`Вы действительно хотите удалить запись?`)) {
        try {
          const response = await storeAuth.axiosInstance.delete(`${server}/${item}/${action}/${id}`);
          console.log(response.status);
          item === 'resume' 
            ? router.push({ name: 'persons', params: { group: 'staffsec' } }) 
            : this.getItem(item, 'view', this.candId);
  
           storeAlert.alertMessage.setAlert('alert-info', `Запись с ID ${id} удалена`);
          } catch (error) {
          console.error(error)
        }
      }
    },

    submitResume: async function (): Promise<void> {
    try {
      const response = await storeAuth.axiosInstance.post(`${server}/resume/${this.action}`, this.form);
      const { message } = response.data;
      storeAlert.alertMessage.setAlert(this.action === "create" 
                                        ? "alert-success" : "alert-info", 
                                        this.action === "create"
                                        ? 'Анкета успешно добавлена' 
                                        : 'Анкета успешно обновлена');
      if (this.action === 'create') {
        this.candId = message
        router.push({ name: 'profile', params: { id: message } });
      } else {
        this.getItem('resume', 'view', this.candId);
      };
        this.cancelEdit();
    } catch (error) {
      console.error(error);
    }
  },
  
    submitFile: async function (
      event: Event, flag: string, idItem: string = '0'
      ): Promise<void> {
  
      const inputElement = event.target as HTMLInputElement;
      if (inputElement && inputElement.files && inputElement.files.length > 0) {
  
        const file = inputElement.files[0];
        const maxSizeInBytes = 1024 * 1024; // 1MB
        
        if (file.size > maxSizeInBytes) {
          storeAlert.alertMessage.setAlert(
            'alert-warning', 
            'File size exceeds the limit. Please select a smaller file.'
          );
          inputElement.value = ''; // Reset the input field
          return;
        }
        const formData = new FormData();
        formData.append('file', inputElement.files[0]);

        try {
          const response = await storeAuth.axiosInstance.post(`${server}/file/${flag}/${idItem}`, formData);
          const { message } = response.data;

          if (flag === 'anketa'){
            storeAlert.alertMessage.setAlert("alert-success", "Данные успешно загружены");
            this.candId = message
            router.push({ name: 'profile', params: { id: message } })
          
          } else if (flag === 'image'){
            this.getImage();
            
          } else {
            storeAlert.alertMessage.setAlert("alert-success", 
                                "Файл или файлы успешно загружен/добавлены");
            this.getItem(flag, 'get', idItem);
          };
        
        } catch (error) {
          console.error(error);
        }
      
      } else {
        storeAlert.alertMessage.setAlert("alert-warning", "Ошибка при загрузке файла");
      }
    },
    
    getImage: async function (): Promise<void> {
      try {
        const response = await storeAuth.axiosInstance.get(`${server}/file/get/${this.candId}`, {responseType: 'blob'});
          this.url = window.URL.createObjectURL(new Blob([response.data]));
      } catch (error) {
        console.error(error);
      }
    },
    
    openForm: function  (item: string, handle: string, idItem = '', form = {}) {
      this.flag = item;
      this.action = handle; 
      if (handle == 'create') {
        this.form = {}
      } else {
        this.itemId = idItem; 
        this.form = form
      };
    },

    cancelEdit: function (): void {
      //clearForm(this.form);
      this.action = '';
      this.flag = '';
    }
  });

  return { dataProfile }
})