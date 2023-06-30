import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const isAuthenticated = ref(false)
  const expirationDate: Ref<null | number> = ref(null)

  function authenticationSuccessHandler(expiration: number) {
    isAuthenticated.value = true
    expirationDate.value = expiration

    console.log(Date.now(), expiration, Date.now() < expiration * 1000)

    localStorage.setItem(
      'user',
      JSON.stringify({
        expiration: expiration * 1000
      })
    )
  }

  function onAppMountAuthenticationVerificationHandler() {
    if (localStorage.getItem('user')) {
      const user = JSON.parse(localStorage.getItem('user')!) as { expiration: number }
      if (user.expiration) {
        const isValid = Date.now() < user.expiration
        console.log('User', user, isValid)
        if (isValid) {
          isAuthenticated.value = true
        } else {
          localStorage.removeItem('user')
        }
      }
    }
  }

  function logoutHandler() {
    isAuthenticated.value = false
    expirationDate.value = null
  }

  return {
    isAuthenticated,
    authenticationSuccessHandler,
    onAppMountAuthenticationVerificationHandler,
    logoutHandler
  }
})
