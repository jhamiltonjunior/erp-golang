package entities

import "time"

type User struct {
	ID       int       `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time `json:"-"`
	Active   uint8     `json:"-"`
}
