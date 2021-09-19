# 利用websocket的chat demo 

使用cobra，所以在命令启动上进行了包装

* 总体
```shell script
$ go run main.go --help  
chat demo

Usage:
  chat [command]

Available Commands:
  client      Start client
  help        Help about any command
  server      Starts a chat server

Flags:
  -h, --help      help for chat
  -v, --version   version for chat

Use "chat [command] --help" for more information about a command.
```
* 服务端

```shell script
$ go run main.go server --help
Starts a chat server

Usage:
  chat server [flags]

Flags:
  -h, --help              help for server
  -l, --listen string     listen address (default ":8000")
  -i, --serverid string   server id (default "demo")

```

* 客户端

```shell script
$ go run main.go client --help
Start client

Usage:
  chat client [flags]

Flags:
  -a, --addr string   server address (default "ws://127.0.0.1:8000")
  -h, --help          help for client
  -u, --user string   server user

```
