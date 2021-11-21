package config

import (
	"testing"

	"go-examples/course/handwriting-web-inf/code_16/framework/provider/config"
	"go-examples/course/handwriting-web-inf/code_28/framework/contract"

	tests "go-examples/course/handwriting-web-inf/code_28/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})

	Convey("test config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{}
		err := configService.Load("database.default", config)
		So(err, ShouldBeNil)
	})

	Convey("test default config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxIdle: 10,
		}
		err := configService.Load("database.read", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxIdle, ShouldEqual, 10)
	})

	Convey("test base config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxOpen: 200,
		}
		err := configService.Load("database", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxOpen, ShouldEqual, 100)
	})

}
