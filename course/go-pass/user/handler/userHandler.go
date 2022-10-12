package handler
import (
	"context"
    "user/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	user "user/proto/user"
)
type UserHandler struct{
     //注意这里的类型是 IUserDataService 接口类型
     UserDataService service.IUserDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *UserHandler) AddUser(ctx context.Context,info *user.UserInfo , rsp *user.Response) error {
	log.Info("Received *user.AddUser request")


	return nil
}

func (e *UserHandler) DeleteUser(ctx context.Context, req *user.UserId, rsp *user.Response) error {
	log.Info("Received *user.DeleteUser request")

	return nil
}

func (e *UserHandler) UpdateUser(ctx context.Context, req *user.UserInfo, rsp *user.Response) error {
	log.Info("Received *user.UpdateUser request")

	return nil
}

func (e *UserHandler) FindUserByID(ctx context.Context, req *user.UserId, rsp *user.UserInfo) error {
	log.Info("Received *user.FindUserByID request")

	return nil
}

func (e *UserHandler) FindAllUser(ctx context.Context, req *user.FindAll, rsp *user.AllUser) error {
	log.Info("Received *user.FindAllUser request")

	return nil
}


