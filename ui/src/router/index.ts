import DailyDisruptionsView from '@/views/DailyDisruptionsView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      // view to track daily disruptions
      path: '/',
      name: 'daily-disruptions',
      meta: { title: 'daily disruptions' },
      component: DailyDisruptionsView,
    },
    {
      // view to manage disruption items
      path: '/manage',
      name: 'manage-disruption-items',
      meta: { title: 'manage disruptions' },
      component: () => import('../views/DisruptionManagerView.vue'),
    },
    {
      // view to show a user's information
      path: '/me',
      name: 'user-me',
      meta: { title: 'whois' },
      component: () => import('../views/UserInfoView.vue'),
    },
    {
      // view add a new disruption item
      path: '/add-disruption-item',
      name: 'add-disruption-item',
      meta: { title: 'add disruption' },
      component: () => import('../views/AddDisruptionItemView.vue'),
    },
  ],
})

export default router
