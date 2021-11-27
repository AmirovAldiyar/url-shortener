package models

type ShortUrl struct {
	ID    int64  `db:"id"`
	Long  string `db:"long"`
	Short string `db:"short"`
}

type JsonShortUrl struct {
	ID    int64  `json:"id"`
	Long  string `json:"long"`
	Short string `json:"short"`
}

type ShortUrlRequest struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}
