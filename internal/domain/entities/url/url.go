package url

import "time"

type URL struct {
	Id             int    `json:"url_id"`
	Description    string `json:"description"`
	DestinationURL string `json:"destination_url"`
	OriginalURL    string `json:"original_url"`
	UserID         int    `json:"_"`
	CreateAt       time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
	Active         uint8
}
