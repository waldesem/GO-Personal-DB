import { defineStore } from 'pinia';

export const alertStore = defineStore('alertStore', () => {

  const alertMessage = ({
    attr: '',
    text: '',

    setAlert: function(attr: string = '', text: string = '') {
      Object.assign(this, {
        attr: attr,
        text: text
      });
      setTimeout(() => {
        this.setAlert();
      }, 10000);
    }
  });
  return { 
    alertMessage 
  }
});

