import { defineStore } from 'pinia';
import { ref } from 'vue';
import { authStore } from '@store/token';
import { server } from '@utilities/utils'

export const messageStore = defineStore('messageStore', () => {

  const storeAuth = authStore();
  
  const messageData = ref({
    messages: [],
    hasPrev: false,
    hasNext: false,
    currentPage: 1
  });

  async function updateMessages(action: string = 'new', page: number = 1): Promise<void> {
    try {
      const response = ['new', 'all', 'read'].includes(action) 
        ? await storeAuth.axiosInstance.get(`${server}/messages/${action}/${page}`)
        : await storeAuth.axiosInstance.delete(`${server}/messages/${action}`);
        
        const [ datas, has_prev, has_next ] = response.data;

        Object.assign(messageData.value, {
          messages: datas,
          hasPrev: has_prev.has_prev,
          hasNext: has_next.has_next,
        });
        
      if (action === 'read') {
        updateMessages('all');
      };

    } catch (error) {
      console.error(error);
    }
  };

  return { 
    messageData,
    updateMessages, 
  }
});

