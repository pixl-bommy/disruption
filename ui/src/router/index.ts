import { createRouter, createWebHistory } from 'vue-router'
import DisruptionManagerView from '../views/DisruptionManagerView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      // view to manage disruption items
      path: '/',
      name: 'home',
      component: DisruptionManagerView,
    },
    {
      // view to show a user's information
      path: '/me',
      name: 'user-me',
      component: () => import('../views/UserInfoView.vue'),
    },
    {
      // view add a new disruption item
      path: '/add-disruption-item',
      name: 'add-disruption-item',
      component: () => import('../views/AddDisruptionItemView.vue'),
    },
  ],
})

export default router
