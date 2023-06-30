<script setup lang="ts">
import type { AxiosInstance } from 'axios'
import { inject, ref } from 'vue'
import { useUserStore } from '../stores/user'
import { useRouter } from 'vue-router'

const $axios = inject('$axios') as AxiosInstance
const userStore = useUserStore()
const router = useRouter()

const email = ref('')
const password = ref('')
const username = ref('')
const isSignup = ref(false)

const formSubmitHandler = async () => {
  let response
  if (isSignup.value) {
    response = await $axios.post('/auth/signup', {
      email: email.value,
      password: password.value,
      username: username.value
    })
  } else {
    response = await $axios.post('/auth/login', {
      email: email.value,
      password: password.value
    })
  }

  // if (response.data['data']) {
  //   userStore.authenticationSuccessHandler(response.data['data']['expirationDate'])
  //   router.push({ name: 'dashboard' })
  // }

  if (response.data['data'] && !isSignup.value) {
    userStore.authenticationSuccessHandler(response.data['data']['expirationDate'])
    router.push({ name: 'dashboard' })
  } else if (response.data['data']) {
    isSignup.value = false
  }

  console.log(email.value, password.value, response)
}
</script>

<template>
  <main>
    <div class="auth-container card">
      <h1>Login</h1>
      <form @submit.prevent="formSubmitHandler">
        <div class="mb-3">
          <!-- <label for="exampleFormControlInput1" class="form-label">Email address</label> -->
          <input
            type="email"
            class="form-control"
            id="email-control"
            placeholder="Email..."
            v-model="email"
          />
        </div>
        <div class="mb-3" v-if="isSignup">
          <!-- <label for="exampleFormControlInput1" class="form-label">Email address</label> -->
          <input
            type="username"
            class="form-control"
            id="username-control"
            placeholder="Username..."
            v-model="username"
          />
        </div>
        <div class="mb-3">
          <!-- <label for="exampleFormControlInput1" class="form-label">Email address</label> -->
          <input
            type="password"
            class="form-control"
            id="password-control"
            placeholder="Password..."
            v-model="password"
          />
        </div>
        <div class="col-auto">
          <button type="submit" class="btn btn-primary mb-3">
            {{ !isSignup ? 'Login' : 'Signup' }}
          </button>
          <button class="btn btn-outline-primary mb-3 ml-2" @click="isSignup = !isSignup">
            {{ !isSignup ? 'Signup' : 'Login' }}
          </button>
        </div>
      </form>
    </div>
  </main>
</template>

<style lang="scss">
main {
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
.auth-container {
  width: 30rem;
  @apply m-auto p-3;
}

h1 {
  @apply text-center;
}
</style>
