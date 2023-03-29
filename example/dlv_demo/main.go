package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	v *viper.Viper
}

func LoadConfigFromYaml(c *config) error {
	c.v = viper.New()
	c.v.SetConfigFile("./config.yaml")

	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	age := c.v.Get("Information.Age")
	name := c.v.Get("Information.Name")
	log.Printf("age:%s, name:%s\n", age, name)

	m := c.v.Sub("information")
	log.Printf("keys:%s, image:%s", m.AllKeys(), m.Get("image"))
	return nil
}

func main() {
	cfg := config{v: viper.New()}
	if err := LoadConfigFromYaml(&cfg); err != nil {
		fmt.Println("Failed read config")
	}
}
