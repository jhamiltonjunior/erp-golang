package user

import "time"

type User struct {
	Id       int
	Name     [255]string
	Email    [255]string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Active   uint8
}
