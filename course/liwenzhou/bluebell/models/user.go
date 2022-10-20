package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type ApiUser struct {
	UserID   int64   `db:"user_id" json:"user_id"`
	Username string  `db:"username" json:"username"`
	Email    *string `db:"email" json:"email"`
	Gender   int32   `db:"gender" json:"gender"`
}
