package core

type Item struct {
	ID        uint64 `db:"id" json:"id"`
	ItemID    uint64 `db:"item_id" json:"item_id"`
	ServiceID uint64 `db:"service_id" json:"service_id"`
	ShortUrl  string `db:"short_url" json:"short_url"`
	DestUrl   string `db:"dest_url" json:"dest_url"`
	IsValid   int8   `db:"is_valid" json:"is_valid"`
	Version   int    `db:"version" json:"version"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}
type ItemReq struct {
	ServiceID uint64 `db:"service_id" json:"service_id"`
	DestUrl   string `db:"dest_url" json:"dest_url"`
	Version   int    `db:"version" json:"version"`
}

func buildItem(i *ItemReq) *Item {
	var r Item
	r.ServiceID = i.ServiceID
	r.DestUrl = i.DestUrl
	r.Version = i.Version
	r.IsValid = 0
	r.ItemID = pero.iID.Add(1)
	return &r
}
