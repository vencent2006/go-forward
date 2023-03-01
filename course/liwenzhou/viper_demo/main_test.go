package main

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"

	"github.com/spf13/viper"
)

func TestReadConf(t *testing.T) {
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

	sub := viper.Sub("mysql")
	fmt.Printf("mysql conf is %#v\n", sub)

	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("conf is %#v\n", C)

	validate := validator.New()
	err = validate.Struct(&C)
	if err != nil {
		fmt.Println("=== error msg ====")
		fmt.Println(err)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("\r\n=========== error field info ====================")
		for _, err := range err.(validator.ValidationErrors) {
			// 列出效验出错字段的信息
			fmt.Println("Namespace: ", err.Namespace())
			fmt.Println("Fild: ", err.Field())
			fmt.Println("StructNamespace: ", err.StructNamespace())
			fmt.Println("StructField: ", err.StructField())
			fmt.Println("Tag: ", err.Tag())
			fmt.Println("ActualTag: ", err.ActualTag())
			fmt.Println("Kind: ", err.Kind())
			fmt.Println("Type: ", err.Type())
			fmt.Println("Value: ", err.Value())
			fmt.Println("Param: ", err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

}
