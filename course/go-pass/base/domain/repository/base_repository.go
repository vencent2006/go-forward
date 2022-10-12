package repository
import (
	"github.com/jinzhu/gorm"
	"base/domain/model"
)
//创建需要实现的接口
type IBaseRepository interface{
    //初始化表
    InitTable() error
    //根据ID查处找数据
    FindBaseByID(int64) (*model.Base, error)
    //创建一条 base 数据
	CreateBase(*model.Base) (int64, error)
    //根据ID删除一条 base 数据
	DeleteBaseByID(int64) error
    //修改更新数据
	UpdateBase(*model.Base) error
    //查找base所有数据
	FindAll()([]model.Base,error)

}
//创建baseRepository
func NewBaseRepository(db *gorm.DB) IBaseRepository  {
	return &BaseRepository{mysqlDb:db}
}

type BaseRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *BaseRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Base{}).Error
}

//根据ID查找Base信息
func (u *BaseRepository)FindBaseByID(baseID int64) (base *model.Base,err error) {
	base = &model.Base{}
	return base, u.mysqlDb.First(base,baseID).Error
}

//创建Base信息
func (u *BaseRepository) CreateBase(base *model.Base) (int64, error) {
	return base.ID, u.mysqlDb.Create(base).Error
}

//根据ID删除Base信息
func (u *BaseRepository) DeleteBaseByID(baseID int64) error {
	return u.mysqlDb.Where("id = ?",baseID).Delete(&model.Base{}).Error
}

//更新Base信息
func (u *BaseRepository) UpdateBase(base *model.Base) error {
	return u.mysqlDb.Model(base).Update(base).Error
}

//获取结果集
func (u *BaseRepository) FindAll()(baseAll []model.Base,err error) {
	return baseAll, u.mysqlDb.Find(&baseAll).Error
}

