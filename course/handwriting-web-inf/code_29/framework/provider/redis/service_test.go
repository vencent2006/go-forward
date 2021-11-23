/**
 * @Author: vincent
 * @Description:
 * @File:  service_test
 * @Version: 1.0.0
 * @Date: 2021/11/21 10:25
 */

package redis

import (
	"context"
	"go-examples/course/handwriting-web-inf/code_29/framework/provider/config"
	"go-examples/course/handwriting-web-inf/code_29/framework/provider/log"
	tests "go-examples/course/handwriting-web-inf/code_29/test"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})

	Convey("test get client", t, func() {
		hadeRedis, err := NewHadeRedis(container)
		So(err, ShouldBeNil)
		service, ok := hadeRedis.(*HadeRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
