/**
 * @Author: vincent
 * @Description:
 * @File:  subject_controller
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:31
 */

package main

import "go-examples/course/handwriting-web-inf/code_05/framework"

func SubjectAddController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
