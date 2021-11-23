/**
 * @Author: vincent
 * @Description:
 * @File:  service_test
 * @Version: 1.0.0
 * @Date: 2021/11/22 17:14
 */

package ssh

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/config"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/log"
	"testing"

	tests "go-examples/course/handwriting-web-inf/code_28/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeSSHService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&SSHProvider{})

	Convey("test get client", t, func() {
		hadeRedis, err := NewHadeSSH(container)
		So(err, ShouldBeNil)
		service, ok := hadeRedis.(*HadeSSH)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		//client, err := service.GetClient()
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		fmt.Println(string(out))
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
