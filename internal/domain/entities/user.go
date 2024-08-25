package entities

import "time"

type UserID int64

type User struct {
	ID       int64     `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time `json:"-"`
	Active   uint8     `json:"-"`
}
