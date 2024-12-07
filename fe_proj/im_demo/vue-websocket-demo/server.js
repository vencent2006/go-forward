const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', function connection(ws) {
  ws.on('message', function incoming(message) {
    console.log('received:', message);
    // Echo the message back to the client
    ws.send(`You said: ${message}`);
  });

  ws.send('Welcome to WebSocket server!');
});

console.log('WebSocket server is running on ws://localhost:8080/');
