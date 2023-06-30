<script setup lang="ts">
import { AxiosInstance } from 'axios'
import { inject, onMounted, ref } from 'vue'
import { useTodoStore } from '../stores/todo'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
const $axios = inject('$axios') as AxiosInstance
const todoStore = useTodoStore()
const userStore = useUserStore()
const todoState = ref('')
const router = useRouter()

onMounted(async () => {
  const response = await $axios.get('/todo')
  console.log('Response: ', response)
  todoStore.fetchTodosHandler(response.data.data)

  const statusResponse = await $axios.get('/todo/status')
  console.log('Response: ', statusResponse)
  todoStore.fetchStatusHandler(statusResponse.data.data)
})

async function onTodoSubmitHandler() {
  console.log(todoState.value)
  try {
    const response = await $axios.post('/todo', {
      body: todoState.value
    })
    console.log('response', response)
    todoStore.todos.push(response.data.data)
  } catch (error) {
    console.log(error)
  }
}

async function onTodoStatusClickHandler(id: string) {
  try {
    const response = await $axios.put(`/todo/status/${id}`)
    console.log('response 2', response)
    todoStore.updateTodoStatusHandler(response.data.data)
  } catch (error) {
    console.error(error)
  }
}

async function onTodoDeleteClickHandler(id: string) {
  console.log('delete', id)
  try {
    const response = await $axios.delete(`/todo/${id}`)
    console.log('response delete', response.data.data, id)

    if (response.data.data === id.toString()) {
      console.log('response delete 2', response)
      todoStore.removeTodoHandler(id)
    }
  } catch (error) {
    console.error(error)
  }
}

async function onLogoutClickHandler() {
  console.log('Logout')
  localStorage.removeItem('user')
  userStore.logoutHandler()
  todoStore.cleanupTodoHandler()
  router.push({ name: 'auth' })
}
</script>

<template>
  <main>
    <div class="todo-container">
      <button class="btn btn-primary" @click="onLogoutClickHandler">Logout</button>
      <form @submit.prevent="onTodoSubmitHandler">
        <div>
          <!-- <label for="exampleFormControlInput1" class="form-label">Email address</label> -->
          <input
            type="text"
            class="form-control"
            id="new-todo-control"
            name="new-todo-control"
            placeholder="Insert new todo..."
            v-model="todoState"
          />
        </div>
        <button type="submit" class="btn btn-primary" :disabled="todoState.length === 0">
          Submit
        </button>
      </form>
      <div class="todo-container-wrapper">
        <div class="todo-container-item" v-for="todo in todoStore.todos" :key="todo.ID">
          <div>
            <button @click="onTodoDeleteClickHandler(todo.ID)">
              <font-awesome-icon :icon="['fas', 'trash']" />
            </button>
            <span>
              {{ todo.body }}
            </span>
          </div>

          <span :status="todo.status" @click="onTodoStatusClickHandler(todo.ID)">
            {{ todo.status }}
          </span>
        </div>
      </div>
    </div>
  </main>
</template>

<style lang="scss">
.todo-container {
  @apply flex flex-col;
  width: 30rem;
  gap: 0.3rem;

  & > form {
    @apply flex;
    margin-bottom: 1rem;

    & > div {
      box-shadow: 0px 0px 8px -6px #000000;
    }

    & > div:first-child {
      flex: 1;
    }
  }

  & > .todo-container-wrapper {
    @apply flex flex-col;
    gap: 0.5rem;

    & > .todo-container-item {
      width: 100%;
      border: 1px solid #dee2e6;
      border-radius: 0.5rem;
      padding: 0.2rem 1rem;
      display: flex;
      justify-content: space-between;
      align-items: center;
      box-shadow: 0px 0px 8px -6px #000000;

      & > div {
        & > button {
          margin-right: 1rem;
          color: initial;
          transition: color 0.3s ease-in;

          &:hover {
            color: #e03616;
          }
        }
      }

      & > span:nth-child(2) {
        border-radius: 2rem;
        padding: 0.2rem 1rem;
        cursor: pointer;
      }

      & > span:nth-child(2)[status|='in progress'] {
        background-color: #ffd275;
      }
      & > span:nth-child(2)[status|='created'] {
        background-color: #56876d;
        color: #f0f0f0;
      }
      & > span:nth-child(2)[status|='done'] {
        background-color: #e03616;
        color: #f0f0f0;
      }
    }
  }
}
</style>
