package entity

type URL struct {
	ID        int64  `json:"id"`
	ShortURL  string `json:"shortUrl"`
	LongURL   string `json:"longUrl"`
	CreatedAt int64  `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}
