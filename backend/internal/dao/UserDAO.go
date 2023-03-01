package dao

import "database/sql"

type User struct {
	ID        string       `json:"id" db:"id"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
