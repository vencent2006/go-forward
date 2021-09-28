/**
 * @Author: vincent
 * @Description:
 * @File:  gorm_test
 * @Version: 1.0.0
 * @Date: 2021/9/27 19:54
 */

package gorm_demo

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BusinessToolOrder struct {
	Id              int64     `json:"id"`
	Uid             int64     `json:"uid"`
	Pid             int       `json:"pid"`
	Oid             int64     `json:"oid"`
	Status          int8      `json:"status"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
	ValidTime       time.Time `json:"valid_time"`
	Price           float64   `json:"price"`
	Extra           string    `json:"extra"`
	Pname           string    `json:"pname"`
	From            int       `json:"from"`
	PayTime         time.Time `json:"pay_time"`
	PayEndTime      time.Time `json:"pay_end_time"`
	ContractNo      string    `json:"contract_no"`
	ContractAddTime int       `json:"contract_add_time"`
	IsOnline        int       `json:"is_online"`
	IsZfb           int       `json:"is_zfb"`
	IsUsed          int       `json:"is_used"`
	ParentId        int64     `json:"parent_id"`
	BuyFrom         string    `json:"buy_from"`
	BuyFid          int       `json:"buy_fid"`
	BuyClient       int       `json:"buy_client"`
	NewPackage      int       `json:"new_package"`
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/eps_package?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)
	// 查询
	var order = new(BusinessToolOrder)
	db.First(order)
	fmt.Printf("%#v\n", order)

	var o BusinessToolOrder
	db.Find(&o, "id=?", 3223418)
	fmt.Printf("%#v\n", o)
}
