/**
 * @Author: vincent
 * @Description:
 * @File:  subject_controller
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:31
 */

package main

import (
	"fmt"
	"go-examples/course/handwriting-web-inf/code_09/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectListController")
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
