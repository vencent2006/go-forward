package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/liwenzhou?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

// Value implement driver.Valuer
func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func queryRowDemo() {
	callStack()
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("sqlx get failed, err:%v\n", err)
		return
	}

	fmt.Printf("user is %#v\n", u)
}

func queryMultiRowDemo() {
	callStack()
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("sqlx.Select failed, err:%v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("db.Exec failed, err:%v\n", err)
		return
	}

	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success, id is %d.\n", id)
}

func updateRowDemo() {
	callStack()
	sqlStr := "update user set age = 111 where id = ?"
	ret, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("ret.RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("update success, affectedRows = %d\n", rowsAffected)
}

func deleteRowDemo() {
	callStack()
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("ret.RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("delete success, affectedRows = %d\n", rowsAffected)
}

func namedExecDemo() {
	_, err := db.NamedExec(`INSERT INTO user (name, age) VALUES (:name, :age)`,
		map[string]interface{}{
			"name": "七米",
			"age":  28,
		})
	if err != nil {
		fmt.Printf("namedExec failed, err:%v\n", err)
	}
}

func namedQueryDemo() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	rows, err := db.NamedQuery(sqlStr,
		map[string]interface{}{
			"name": "七米",
		})
	if err != nil {
		fmt.Printf("namedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{Name: "七米"}
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

}

func transactionDemo() (err error) {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("Beginx failed, err:%v\n", tx)
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// 发生了panic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit done")
		}
	}()

	// sqlStr1
	sqlStr1 := `UPDATE user set age=20 where id = ?`
	ret, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	// sqlStr2
	sqlStr2 := `UPDATE user set age=50 where id = ?`
	ret, err = tx.Exec(sqlStr2, 2)
	if err != nil {
		return err
	}
	affected, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("exec sqlStr2 failed")
	}

	return err
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*user) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
		fmt.Println("valueStrings:", valueStrings)
		fmt.Println("valueArgs:", valueArgs)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	fmt.Println("stmt:", stmt)
	fmt.Println("valueArgs:", valueArgs)
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func batchInsertDemos() {
	u1 := user{Name: "七米", Age: 18}
	u2 := user{Name: "q1mi", Age: 28}
	u3 := user{Name: "小王子", Age: 38}
	users := []*user{&u1, &u2, &u3}
	err := batchInsertUsers3(users)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func batchInsertUsers3(users []*user) error {
	_, err := db.NamedExec("INSERT INTO user(name, age) VALUES (:name, :age)", users)
	return err
}

func QueryByIDs(ids []int) (users []user, err error) {
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}

	query = db.Rebind(query)

	err = db.Select(&users, query, args...)

	return
}

func QueryByIDsDemo() {
	ids := []int{4, 1, 2, 5}
	fmt.Println(reflect.TypeOf(ids))
	users, err := QueryByIDs(ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []user, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)",
		ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func QueryAndOrderByIDsDemo() {
	ids := []int{4, 1, 2, 5}
	fmt.Println(reflect.TypeOf(ids))
	users, err := QueryAndOrderByIDs(ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

func callStack() {
	caller, _, _, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("********** in %s\n", runtime.FuncForPC(caller).Name())
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed, err:%v\n", err)
		return
	}
	fmt.Println("init DB success...")

	//queryRowDemo()
	//queryMultiRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
	//namedExecDemo()
	//namedQueryDemo()
	//_ = transactionDemo()
	//batchInsertDemos()
	//QueryByIDsDemo()
	QueryAndOrderByIDsDemo()
}
