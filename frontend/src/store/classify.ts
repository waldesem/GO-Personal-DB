import axios from 'axios';
import { ref } from 'vue';
import { defineStore } from 'pinia'
import { server } from '@utilities/utils';


export const classifyStore = defineStore('classifyStore', () => {

  const classData = ref({
    status: <Record<string, any>>({}),
    regions: <Record<string, any>>([{}]),
    conclusion: <Record<string, any>>({}),
    category: <Record<string, any>>({}),
    groups: <Record<string, any>>({}),
    roles: <Record<string, any>>({}),
    tables: Array<string>(),
    getClasses: async function (): Promise<void> {
      try {
        const response = await axios.get(`${server}/classes`);
        const {
          category,
          conclusion,
          role,
          group,
          status,
          region,
          tables
        } = response.data;

        Object.assign(classData.value, {
          category: category,
          conclusion: conclusion,
          roles: role,
          groups: group,
          status: status.reduce((acc: { [x: string]: any; }, 
                                item: { id: string | number; status: any; }) => {
            acc[item.id] = item.status;
            return acc;
          }),
          regions: region.reduce((acc: { [x: string]: any; },
                                 item: { id: string | number; region: any; }) => {
            acc[item.id] = item.region;
            return acc;
          }, {} as { [key: string]: string }),
          tables: tables
        })
      } catch (error) {
        console.error(error)
      }
    }
  });
  return {
    classData,
  }
});
