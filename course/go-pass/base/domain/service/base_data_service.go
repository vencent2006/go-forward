package service

import (
	"base/domain/model"
	"base/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IBaseDataService interface {
	AddBase(*model.Base) (int64 , error)
	DeleteBase(int64) error
	UpdateBase(*model.Base) error
	FindBaseByID(int64) (*model.Base, error)
	FindAllBase() ([]model.Base, error)
}


//创建
//注意：返回值 IBaseDataService 接口类型
func NewBaseDataService(baseRepository repository.IBaseRepository,clientSet *kubernetes.Clientset) IBaseDataService{
	return &BaseDataService{ BaseRepository:baseRepository, K8sClientSet: clientSet,deployment:&v1.Deployment{}}
}

type BaseDataService struct {
    //注意：这里是 IBaseRepository 类型
	BaseRepository repository.IBaseRepository
	K8sClientSet  *kubernetes.Clientset
	deployment  *v1.Deployment
}


//插入
func (u *BaseDataService) AddBase(base *model.Base) (int64 ,error) {
	 return u.BaseRepository.CreateBase(base)
}

//删除
func (u *BaseDataService) DeleteBase(baseID int64) error {
	return u.BaseRepository.DeleteBaseByID(baseID)
}

//更新
func (u *BaseDataService) UpdateBase(base *model.Base) error {
	return u.BaseRepository.UpdateBase(base)
}

//查找
func (u *BaseDataService) FindBaseByID(baseID int64) (*model.Base, error) {
	return u.BaseRepository.FindBaseByID(baseID)
}

//查找
func (u *BaseDataService) FindAllBase() ([]model.Base, error) {
	return u.BaseRepository.FindAll()
}

