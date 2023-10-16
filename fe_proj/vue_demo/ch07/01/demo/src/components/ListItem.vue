<template>
  <li>{{ msg }}</li>
</template>

<script>
export default {
  name: 'ListItem',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
