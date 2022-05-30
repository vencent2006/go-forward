package services

import (
	"go-examples/course/gateway/go_gateway_demo/dao"
	"go-examples/course/gateway/go_gateway_demo/dto"
	"go-examples/course/gateway/go_gateway_demo/middleware"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
}

func (as *ApiService) AddUser(c *gin.Context, input *dto.AddUserInput) error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return err
	}
	user := &dao.User{
		Name:  input.Name,
		Sex:   input.Sex,
		Age:   input.Age,
		Birth: input.Birth,
		Addr:  input.Addr,
	}
	if err := user.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2002, err)
		return err
	}
	return nil
}
