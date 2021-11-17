package config

import (
	"path/filepath"
	"testing"

	"go-examples/course/handwriting-web-inf/code_25/framework/contract"
	tests "go-examples/course/handwriting-web-inf/code_25/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_GetInt(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewHadeConfig(folder, map[string]string{}, contract.EnvDevelopment)
		So(err, ShouldBeNil)
		conf := serv.(*HadeConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
