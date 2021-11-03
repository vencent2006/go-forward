/**
 * @Author: vincent
 * @Description:
 * @File:  provider_test
 * @Version: 1.0.0
 * @Date: 2021/11/1 14:35
 */

package env

import (
	"go-examples/course/handwriting-web-inf/code_18/framework"
	"go-examples/course/handwriting-web-inf/code_18/framework/contract"
	"go-examples/course/handwriting-web-inf/code_18/framework/provider/app"
	"testing"

	tests "go-examples/course/handwriting-web-inf/code_18/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeEnvProvider(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		sp := &app.HadeAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &HadeEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
