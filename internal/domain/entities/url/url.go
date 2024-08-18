package url

import "time"

type URL struct {
	Id             int
	DestinationURL string
	OriginalURL    string
	UserID         int
	CreateAt       time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
	Active         uint8
}
