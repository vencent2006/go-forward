const { log } = require('console')
var http = require('http')

http
  .createServer(function (req, res) {
    res.writeHead(200, { 'Content-Type': 'text/plain' })
    res.end('Hello world\n')
  })
  .listen(8888)

console.log('server is running at http://127.0.0.1:8888')
