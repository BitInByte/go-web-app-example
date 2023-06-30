import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import AuthView from '../views/AuthView.vue'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'auth',
      component: AuthView
    },
    {
      path: '/',
      name: 'dashboard',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/DashboardView.vue')
    }
  ]
})

router.beforeEach(async (to, from) => {
  const userStore = useUserStore()

  if (to.name == 'dashboard' && !userStore.isAuthenticated) {
    return { name: 'auth' }
  } else if (to.name == 'auth' && userStore.isAuthenticated) {
    return { name: 'dashboard' }
  }
})

export default router
