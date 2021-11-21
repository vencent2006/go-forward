/**
 * @Author: vincent
 * @Description:
 * @File:  model
 * @Version: 1.0.0
 * @Date: 2021/11/16 18:36
 */

package demo

import (
	"database/sql"
	"time"
)

type UserModel struct {
	UserId int
	Name   string
	Age    int
}

// User is gorm model
type User struct {
	ID           uint
	Name         string
	Email        *string // todo: 为啥是*string
	Age          uint8
	Birthday     *time.Time // todo: 为啥是*time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
