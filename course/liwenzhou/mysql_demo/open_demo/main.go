package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 自动执行init
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/eps_package"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数

	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("connect db succ")
}
