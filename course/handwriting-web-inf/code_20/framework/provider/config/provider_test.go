package config

import (
	"testing"

	"go-examples/course/handwriting-web-inf/code_20/framework"
	"go-examples/course/handwriting-web-inf/code_20/framework/contract"
	"go-examples/course/handwriting-web-inf/code_20/framework/provider/app"
	"go-examples/course/handwriting-web-inf/code_20/framework/provider/env"

	tests "go-examples/course/handwriting-web-inf/code_20/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_Normal(t *testing.T) {
	Convey("test hade config normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		c.Bind(&app.HadeAppProvider{BaseFolder: basePath})
		c.Bind(&env.HadeEnvProvider{})

		err := c.Bind(&HadeConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.mysql.hostname"), ShouldEqual, "127.0.0.1")
		So(conf.GetInt("database.mysql.timeout"), ShouldEqual, 1)
		So(conf.GetFloat64("database.mysql.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.mysql")
		So(maps, ShouldContainKey, "hostname")
		So(maps["timeout"], ShouldEqual, 1)

		maps2 := conf.GetStringMapString("databse.mysql")
		So(maps2["timeout"], ShouldEqual, "")

		type Mysql struct {
			Hostname string
			Username string
		}
		ms := &Mysql{}
		err = conf.Load("database.mysql", ms)
		Println(ms)
		So(err, ShouldBeNil)
		So(ms.Hostname, ShouldEqual, "127.0.0.1")
	})
}
