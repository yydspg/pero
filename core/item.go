package core

type Item struct {
	ID        int64  `db:"id" json:"id"`
	ItemID    int64  `db:"item_id" json:"item_id"`
	ServiceID int64  `db:"service_id" json:"service_id"`
	ShortUrl  string `db:"short_url" json:"short_url"`
	DestUrl   string `db:"dest_url" json:"dest_url"`
	IsValid   int8   `db:"is_valid" json:"is_valid"`
	Version   int    `db:"version" json:"version"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}
