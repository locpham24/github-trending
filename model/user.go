package model

import "time"

type User struct {
	UserId    string    `db:"user_id, omitempty"`
	FullName  string    `db:"full_name, omitempty"`
	Email     string    `db:"email, omitempty"`
	Password  string    `json:"password,omitempty" db:"password, omitempty"`
	Role      string    `json:"-" db:"role, omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"token,omitempty"`
}
