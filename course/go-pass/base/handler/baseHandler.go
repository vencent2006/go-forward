package handler
import (
	"context"
    "base/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	base "base/proto/base"
)
type BaseHandler struct{
     //注意这里的类型是 IBaseDataService 接口类型
     BaseDataService service.IBaseDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *BaseHandler) AddBase(ctx context.Context,info *base.BaseInfo , rsp *base.Response) error {
	log.Info("Received *base.AddBase request")


	return nil
}

func (e *BaseHandler) DeleteBase(ctx context.Context, req *base.BaseId, rsp *base.Response) error {
	log.Info("Received *base.DeleteBase request")

	return nil
}

func (e *BaseHandler) UpdateBase(ctx context.Context, req *base.BaseInfo, rsp *base.Response) error {
	log.Info("Received *base.UpdateBase request")

	return nil
}

func (e *BaseHandler) FindBaseByID(ctx context.Context, req *base.BaseId, rsp *base.BaseInfo) error {
	log.Info("Received *base.FindBaseByID request")

	return nil
}

func (e *BaseHandler) FindAllBase(ctx context.Context, req *base.FindAll, rsp *base.AllBase) error {
	log.Info("Received *base.FindAllBase request")

	return nil
}


