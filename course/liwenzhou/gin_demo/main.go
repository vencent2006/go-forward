package main

import (
	"flag"
	"fmt"
	"liwenzhou/gin_demo/global"
	"liwenzhou/gin_demo/internal/middleware"
	tracer "liwenzhou/gin_demo/pkg"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/gin-gonic/gin"
)

var port string
var db *gorm.DB

const (
	constServiceName   = "blog-service"
	constAgentHostPort = "127.0.0.1:6831"
)

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(constServiceName, constAgentHostPort)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

// Post 为内存对齐对字段顺序进行了调整
type Post struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	PostId      int64     `json:"post_id" gorm:"column:post_id"`
	AuthorId    int64     `json:"author_id" gorm:"column:author_id"`
	CommunityId int64     `json:"community_id" gorm:"column:community_id"`
	Status      int32     `json:"status" gorm:"column:status"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func setupGorm() (err error) {
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	otgorm.AddGormCallbacks(db)
	return
}

func main() {
	// 1. parse flag
	flag.StringVar(&port, "p", "8080", "addr port")
	flag.Parse()

	// 2. set up
	var err error
	err = setupTracer() // setup tracer with jaeger
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupGorm()
	if err != nil {
		log.Fatalf("init.setupGorm err: %v", err)
	}
	defer db.Close()

	// 3. gin instance
	r := gin.Default()

	// 4. router and handler
	r.Use(middleware.Tracing())
	r.GET("/hello", func(c *gin.Context) {
		// 查询
		dbEngine := otgorm.WithContext(c.Request.Context(), db)
		var u = new(Post)
		dbEngine.First(u)
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, u)
	})

	r.Run(":" + port)
}
