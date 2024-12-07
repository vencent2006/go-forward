<template>
  <div id="app">
    <h1>Vue3 WebSocket Demo</h1>
    <input v-model="message" placeholder="Enter your message" @keyup.enter="sendMessage" />
    <button @click="sendMessage">Send</button>
    <p>Received: {{ receivedMessage }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: '',
      receivedMessage: '',
      socket: null,
    };
  },
  created() {
    this.socket = new WebSocket('ws://localhost:8080');

    this.socket.onopen = () => {
      console.log('WebSocket connection opened');
    };

    this.socket.onmessage = (event) => {
      this.receivedMessage = event.data;
    };

    this.socket.onclose = () => {
      console.log('WebSocket connection closed');
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };
  },
  methods: {
    sendMessage() {
      if (this.message.trim()) {
        this.socket.send(this.message);
        this.message = '';
      }
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
}
input {
  padding: 10px;
  margin: 10px 0;
  width: 200px;
}
button {
  padding: 10px 20px;
  margin: 10px;
}
</style>
