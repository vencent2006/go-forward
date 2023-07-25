package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
-- zero.vip_user_ai_blog definition

CREATE TABLE `vip_user_ai_blog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` bigint(20) NOT NULL COMMENT '用户id',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '类型:1用户主动2定时推送',
  `image_url` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户上传图片',
  `text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户上传文案',
  `machine_data` varchar(2048) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '机器学习侧返回结果',
  `machine_ret` varchar(2048) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '结果',
  `machinei_time` int(11) NOT NULL DEFAULT '0' COMMENT '请求时长',
  `secure` varchar(2048) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '安全侧返回结果',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1是 0否',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_type_uid` (`type`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ai发博文';
*/

type VipUserAiBlog struct {
	Id          int64  `json:"id"`
	Uid         int64  `json:"uid"`
	MachineData string `json:"machine_data"`
	MachineRet  string `json:"machine_ret"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "test:123456@(10.41.41.203:3306)/zero?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
}

func main() {
	defer db.Close()
	r := gin.Default()
	r.GET("/blog", getBlog)
	r.Run(":8081")
}

func getBlog(c *gin.Context) {
	blogId := c.Query("id")
	var blog VipUserAiBlog
	db.Find(&blog, "id=?", blogId)
	fmt.Printf("%#v\n", blog)

	c.JSON(200, blog)
}
