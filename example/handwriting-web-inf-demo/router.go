/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/10/24 16:24
 */

package main

import "go-examples/example/handwriting-web-inf-demo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
