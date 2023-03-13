package dao

import "database/sql"

type Todo struct {
	ID        string       `json:"id" db:"id"`
	UID       string       `json:"uid" db:"uid"`
	Todo      string       `json:"todo" db:"todo"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
