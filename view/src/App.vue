<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { inject, onMounted } from 'vue'
import type { AxiosInstance } from 'axios'
import { useUserStore } from './stores/user'
const userStore = useUserStore()

// const $axios = inject('$axios') as AxiosInstance
//
// onMounted(async () => {
//   console.log('url', window.location.origin)
//   const result = await $axios.post('/auth/login', {
//     email: 'test@test.com',
//     password: 'test123'
//   })
//   console.log('Result: ', result)
// })
onMounted(() => {
  userStore.onAppMountAuthenticationVerificationHandler()
})
</script>

<template>
  <RouterView />
</template>
