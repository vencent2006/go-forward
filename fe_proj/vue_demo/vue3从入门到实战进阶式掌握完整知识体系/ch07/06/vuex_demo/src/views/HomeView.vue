<template>
  <div class="home">
    <h1>{{ name }}</h1>
  </div>
</template>

<script>
// @ is an alias to /src
import { toRefs } from 'vue';
import { useStore } from 'vuex';
export default {
  name: 'HomeView',
  setup() {
    const store = useStore()
    const { name } = toRefs(store.state)
    return { name }
  }
}
</script>
