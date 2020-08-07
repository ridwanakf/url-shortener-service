package entity

import "time"

type URL struct {
	ID        int64     `json:"id" db:"id"`
	ShortURL  string    `json:"shortUrl" db:"short_url"`
	LongURL   string    `json:"longUrl" db:"long_url"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	ExpireAt  time.Time `json:"expireAt" db:"expire_at"`
	CreatedBy string    `json:"createdBy" db:"created_by"`
}
