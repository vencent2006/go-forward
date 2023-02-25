package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type conf struct {
	Port        int    `mapstructrure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	DbName string `mapstructure:"dbname"`
	Port   string `mapstructure:"port"`
}

var C conf

func main() {
	viper.SetDefault("fileDir", "/")
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	// 实时监控配置文件的变化
	viper.WatchConfig()
	// 当配置变化之后调用的一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		sub := viper.Sub("mysql")
		fmt.Printf("mysql conf is %#v\n", sub)

		err := viper.Unmarshal(&C)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("conf is %#v\n", C)

		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run()
}
