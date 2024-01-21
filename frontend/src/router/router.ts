import { createRouter, createWebHistory } from 'vue-router'
import { server } from '@utilities/utils';
import axios from 'axios';

const router = createRouter({
  routes: [
    {
      path: '/',
      component: () => import('@/App.vue')
    },
    {
      path: '/login/auth',
      name: 'login',
      component: () => import('@components/pages/LoginPage.vue')
    },
    {
      path: '/:group',
      name: 'group',
      component: () => import('@components/PagesVue.vue'),
      children: [
        {
          path: 'persons',
          name: 'persons',
          component: () => import('@components/pages/PersonPage.vue'),
        },
        {
          path: 'resume',
          name: 'resume',
          component: () => import('@components/pages/ResumePage.vue')
        },
        {
          path: 'profile/:id',
          name: 'profile',
          component: () => import('@components/pages/ProfilePage.vue')
        },
        {
          path: 'information',
          name: 'information',
          component: () => import('@components/pages/InfoPage.vue')
        },
        {
          path: 'users',
          name: 'users',
          component: () => import('@components/pages/AdminPage.vue')
        },
        {
          path: 'user/:id',
          name: 'user',
          component: () => import('@components/pages/UserPage.vue')
        },
        {
          path: 'table',
          name: 'table',
          component: () => import('@components/pages/TablesPage.vue')
        },
        {
          path: 'contacts',
          name: 'contacts',
          component: () => import('@components/pages/ContactPage.vue')
        },
        {
          path: 'manager',
          name: 'manager',
          component: () => import('@components/pages/FilePage.vue'),
          props: route => ({ path: route.query.path })
        },
        {
          path: 'messages',
          name: 'messages',
          component: () => import('@components/pages/MessagePage.vue')
        }
      ]
    },
    {
      path: '/:pathMatch(.*)*',  
      name: '404', 
      component: () => import('@components/pages/NotFound.vue')
    }
  ],
  history: createWebHistory()
});

router.beforeEach(async (to, _from, next) => {
    
  const refresh_token = localStorage.getItem('refresh_token');
  const access_token = localStorage.getItem('access_token');

  if (to.name !== 'login') {
    if (refresh_token) {

      const expiry_refresh = (JSON.parse(atob(refresh_token.split('.')[1]))).exp;

      if (Math.floor((new Date).getTime() / 1000) >= expiry_refresh) {
        next({ name: 'login' });
        
      } else {

        if (access_token) {
          const expiry_access = (JSON.parse(atob(access_token.split('.')[1]))).exp;

          if (Math.floor((new Date).getTime() / 1000) >= expiry_access) {
            const response = await axios.post(`${server}/refresh`, null, {
              headers: {'Authorization': `Bearer ${localStorage.getItem('refresh_token')}`}
            });
            const { access_token } = response.data;
            
            if (access_token){
              localStorage.setItem('access_token', access_token);
              next();

            } else {
              router.push({ name: 'login' });
            };

          } else {
            next();
          };
          
        } else {
          next({ name: 'login' });
        };
      };
      
    } else {
      next({ name: 'login' });
    };
    
  } else {
    next();
  }
});

export default router;