package main

import (
	"os"
	"os/signal"
	"syscall"

	"go-examples/course/gateway/go_gateway_demo/router"

	"github.com/e421083458/golang_common/lib"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
