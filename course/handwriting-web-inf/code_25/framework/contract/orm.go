/**
 * @Author: vincent
 * @Description:
 * @File:  orm
 * @Version: 1.0.0
 * @Date: 2021/11/16 14:02
 */

package contract

import (
	"go-examples/course/handwriting-web-inf/code_25/framework"
	"net"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

// ORMKey 代表 ORM的服务
const ORMKey = "hade:orm"

// ORMService 表示传入的参数
type ORMService interface {
	GetDB(option ...DBOption) (*gorm.DB, error)
}

// DBOption 代表初始化时候的选项
type DBOption func(container framework.Container, config *DBConfig) error

// DBConfig 代表数据库连接的所有配置
type DBConfig struct {
	// 以下配置关于dsn
	WriteTimeout string `yaml:"writeTimeout"` // 写超时时间
	Loc          string `yaml:"loc"`          // 时区
	Port         int    `yaml:"port"`         // 端口
	ReadTimeout  string `yaml:"readTimeout"`  // 读超时时间
	Charset      string `yaml:"charset"`      // 字符集
	ParseTime    bool   `yaml:"parseTime"`    // 是否解析时间
	Protocol     string `yaml:"protocol"`     // 传输协议
	Dsn          string `yaml:"dsn"`          // 直接传递dsn，如果传递了，其他关于dsn的配置均无效
	Database     string `yaml:"database"`     // 数据库
	Collation    string `yaml:"collation"`    // 字符序
	Timeout      string `yaml:"timeout"`      // 连接超时时间
	Username     string `yaml:"username"`     // 用户名
	Password     string `yaml:"password"`     // 密码
	Driver       string `yaml:"driver"`       // 驱动
	Host         string `yaml:"host"`         // 数据库地址

	// 以下配置关于连接池
	ConnMaxIdle     int    `yaml:"connMaxIdle"`     // 最大空闲连接数
	ConnMaxOpen     int    `yaml:"connMaxOpen"`     // 最大连接数
	ConnMaxLifetime string `yaml:"connMaxLifetime"` // 连接最大生命周期
	ConnMaxIdletime string `yaml:"connMaxIdletime"` // 空闲最大生命周期

	// 以下配置关于gorm
	*gorm.Config // 集成gorm的配置
}

// FormatDsn 生成dsn
func (conf *DBConfig) FormatDsn() (string, error) {
	port := strconv.Itoa(conf.Port)
	timeout, err := time.ParseDuration(conf.Timeout)
	if err != nil {
		return "", err
	}
	readTimeout, err := time.ParseDuration(conf.ReadTimeout)
	if err != nil {
		return "", err
	}
	writeTimeout, err := time.ParseDuration(conf.WriteTimeout)
	if err != nil {
		return "", err
	}
	location, err := time.LoadLocation(conf.Loc)
	if err != nil {
		return "", err
	}

	driverConf := &mysql.Config{
		User:                 conf.Username,
		Passwd:               conf.Password,
		Net:                  conf.Protocol,
		Addr:                 net.JoinHostPort(conf.Host, port),
		DBName:               conf.Database,
		Collation:            conf.Collation,
		Loc:                  location,
		Timeout:              timeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		ParseTime:            conf.ParseTime,
		AllowNativePasswords: true, // todo：vincent find， 不加上这个，授权会有问题
	}
	return driverConf.FormatDSN(), nil
}
