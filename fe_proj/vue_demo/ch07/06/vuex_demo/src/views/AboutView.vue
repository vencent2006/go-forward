<template>
  <div class="about">
    <h1 @click="handleClick">This is an about page</h1>
    <h1>{{ name }}</h1>
  </div>
</template>

<script>
import { toRefs } from 'vue'
import { useStore } from 'vuex'
export default {
  name: 'AboutView',
  setup() {
    const store = useStore()
    const { name } = toRefs(store.state)
    const handleClick = () => {
      store.dispatch('getData')
    }
    return { name, handleClick }
  }
}
</script>