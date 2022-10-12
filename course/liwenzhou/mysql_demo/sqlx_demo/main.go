package main

import (
	"errors"
	"fmt"
	"runtime"

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
	namedQueryDemo()
	_ = transactionDemo()
}
