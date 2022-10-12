package service

import (
	"user/domain/model"
	"user/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IUserDataService interface {
	AddUser(*model.User) (int64 , error)
	DeleteUser(int64) error
	UpdateUser(*model.User) error
	FindUserByID(int64) (*model.User, error)
	FindAllUser() ([]model.User, error)
}


//创建
//注意：返回值 IUserDataService 接口类型
func NewUserDataService(userRepository repository.IUserRepository,clientSet *kubernetes.Clientset) IUserDataService{
	return &UserDataService{ UserRepository:userRepository, K8sClientSet: clientSet,deployment:&v1.Deployment{}}
}

type UserDataService struct {
    //注意：这里是 IUserRepository 类型
	UserRepository repository.IUserRepository
	K8sClientSet  *kubernetes.Clientset
	deployment  *v1.Deployment
}


//插入
func (u *UserDataService) AddUser(user *model.User) (int64 ,error) {
	 return u.UserRepository.CreateUser(user)
}

//删除
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

//更新
func (u *UserDataService) UpdateUser(user *model.User) error {
	return u.UserRepository.UpdateUser(user)
}

//查找
func (u *UserDataService) FindUserByID(userID int64) (*model.User, error) {
	return u.UserRepository.FindUserByID(userID)
}

//查找
func (u *UserDataService) FindAllUser() ([]model.User, error) {
	return u.UserRepository.FindAll()
}

