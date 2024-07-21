package core

import "time"

type Item struct {
	ID        uint64    `db:"id" json:"id"`
	ItemID    uint64    `db:"item_id" json:"item_id"`
	ServiceID uint64    `db:"service_id" json:"service_id"`
	ShortUrl  string    `db:"short_url" json:"short_url"`
	DestUrl   string    `db:"dest_url" json:"dest_url"`
	IsValid   int8      `db:"is_valid" json:"is_valid"`
	Version   int       `db:"version" json:"version"`
	CreateAt  time.Time `db:"create_at" json:"create_at" gorm:"column:create_at;type:timestamp"`
	UpdateAt  time.Time `db:"update_at" json:"update_at" gorm:"column:update_at;type:timestamp"`
}
type ItemReq struct {
	ServiceID uint64 `db:"service_id" json:"service_id"`
	DestUrl   string `db:"dest_url" json:"dest_url"`
	Version   int    `db:"version" json:"version"`
}
type ItemUpdateReq struct {
	ServiceID uint64 `db:"service_id" json:"service_id"`
	ItemID    uint64 `db:"item_id" json:"item_id"`
	Version   int    `db:"version" json:"version"`
	IsValid   int8   `db:"is_valid" json:"is_valid"`
}

func buildItemFromCreate(i *ItemReq) *Item {
	var r Item
	r.ServiceID = i.ServiceID
	r.DestUrl = i.DestUrl
	r.Version = i.Version
	r.IsValid = 0
	r.ItemID = getID(r.DestUrl)
	r.ShortUrl = getLink(i.DestUrl)
	r.CreateAt = time.Now()
	r.UpdateAt = time.Now()
	return &r
}
func buildItemFromUpdate(i *ItemUpdateReq) *Item {
	var r Item
	r.ServiceID = i.ServiceID
	r.ItemID = i.ItemID
	r.Version = i.Version
	r.IsValid = i.IsValid
	r.UpdateAt = time.Now()
	return &r
}
