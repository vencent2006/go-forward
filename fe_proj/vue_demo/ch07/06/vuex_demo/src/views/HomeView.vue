<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    {{ myName }}
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'HomeView',
  // 计算属性
  computed: {
    myName() {
      return this.$store.state.name
    }
  }
}
</script>
