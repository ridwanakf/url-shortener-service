package entity

import "time"

type URL struct {
	ID        int64     `json:"id"`
	ShortURL  string    `json:"shortUrl"`
	LongURL   string    `json:"longUrl"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}
