/**
 * @Author: vincent
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/9/20 15:18
 */

package main

import "go-examples/course/handwriting-web-inf/code_02/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
