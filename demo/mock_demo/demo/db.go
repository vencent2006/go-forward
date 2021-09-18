/**
 * @Author: vincent
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2021/8/17 16:27
 */

package main

// db.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
