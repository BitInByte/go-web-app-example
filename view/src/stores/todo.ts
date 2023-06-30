import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'

export const useTodoStore = defineStore('todo', () => {
  const todos: Ref<{ body: string; status: string; ID: string }[]> = ref([])
  const status: Ref<string[]> = ref([])

  function fetchTodosHandler(fetchedTodos: [any]) {
    todos.value = fetchedTodos
  }

  function fetchStatusHandler(fetchedStatus: [string]) {
    status.value = fetchedStatus
  }

  function updateTodoStatusHandler(todo: { body: string; status: string; ID: string }) {
    const index = todos.value.findIndex((item) => item.ID === todo.ID)

    if (index > -1) {
      console.log('response 3', todos.value[index].status, todo.status)
      todos.value[index].status = todo.status
    }
  }

  function removeTodoHandler(id: string) {
    todos.value = todos.value.filter((item) => item.ID !== id)
  }

  function cleanupTodoHandler() {
    todos.value = []
    status.value = []
  }

  return {
    todos,
    fetchTodosHandler,
    status,
    fetchStatusHandler,
    updateTodoStatusHandler,
    removeTodoHandler,
    cleanupTodoHandler
  }
})
