/**
 * @Author: vincent
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2021/10/29 07:24
 */

package demo

import (
	demoService "go-examples/course/handwriting-web-inf/code_27/app/provider/demo"
	"go-examples/course/handwriting-web-inf/code_27/framework/contract"
	"go-examples/course/handwriting-web-inf/code_27/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	// todo: 这个Bind没有效果呢, 但是在main.go使用container.Bind就好使
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.GET("/demo/demo3", api.Demo3)
	r.GET("/demo/orm", api.DemoOrm)
	r.GET("/demo/cache/redis", api.DemoRedis)
	r.GET("/demo/cache/cache", api.DemoCache)

	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取数据库密码
// @Description 获取数据库密码，不用分页
// @Produce  json
// @Tags demo
// @Success 200 {string} json "{"Password":password}"
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	// 获取password
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	res := map[string]interface{}{
		"Password": password,
	}

	logService := c.MustMake(contract.LogKey).(contract.Log)
	logService.Debug(c, "hello", nil)
	logService.Fatal(c, "fatal", nil)
	// 打印出来
	c.JSON(200, res)
}

// Demo2  for godoc
// @Summary 获取app的baseFolder
// @Description 获取app的baseFolder,不进行分页
// @Produce  json
// @Tags demo
// @Success 200 {string} baseFolder
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	baseFolder := appService.BaseFolder()
	c.JSON(200, baseFolder)
}

// Demo godoc
// @Summary 获取所有学生
// @Description 获取所有学生
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo3 [get]
func (api *DemoApi) Demo3(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}
