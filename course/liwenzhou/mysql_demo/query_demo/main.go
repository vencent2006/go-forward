package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql" // 自动执行init
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/liwenzhou"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(100) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数

	return
}

type user struct {
	id   int
	name string
	age  int
}

// 单行
func queryRowDemo() {
	callStack()
	caller, _, _, ok := runtime.Caller(0)
	if ok {
		fmt.Println(runtime.FuncForPC(caller).Name())
	}
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	// 非常重要： 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多行
func queryMultiRowDemo() {
	callStack()

	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要： 关闭rows释放持有的数据库连接
	defer rows.Close()

	// 循环取出结果中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRowDemo() {
	callStack()

	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	lastInsertId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success, lastInsertId: %d\n", lastInsertId)
}

// 更新语句
func updateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("update success, affected rows:%d\n", rowsAffected)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("delete success, affected row number:%d\n", rowsAffected)
}

func callStack() {
	caller, _, _, ok := runtime.Caller(1)
	if ok {
		fmt.Println("-----------------      ", runtime.FuncForPC(caller).Name())
	}
}

// 预处理查询示例
func prepareQueryDemo() {
	callStack()
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("stmt Query failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	// 循环读取结果中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("rows.Scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	callStack()

	sqlStr := "insert into user(name, age) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	fmt.Println("insert success")

}

func main() {
	if err := initDB(); err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("connect db succ")

	//queryRowDemo()
	//queryMultiRowDemo()
	////insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()

	prepareQueryDemo()
	prepareInsertDemo()
	prepareQueryDemo()

	fmt.Println("查询结束了...")
}
