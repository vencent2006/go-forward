package main

import (
	"fmt"
	"go-examples/course/go-redis-rewrite/go-redis/config"
	"go-examples/course/go-redis-rewrite/go-redis/lib/logger"
	"go-examples/course/go-redis-rewrite/go-redis/tcp"
	"os"
)

const configFile string = "redis.conf"

var defautProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})

	if fileExists(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defautProperties
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{
			Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
		},
		tcp.MakeHandler(),
	)
	if err != nil {
		logger.Error(err)
	}
}
