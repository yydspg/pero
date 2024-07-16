package core

import (
	"gorm.io/gorm"
)

type ItemDB struct {
	db *gorm.DB
}

func newItemDB() *ItemDB {
	return &ItemDB{
		getPeroDB().db,
	}
}
func (i *ItemDB) insert(it *Item) {
	i.db.Create(it)
}
func (i *ItemDB) delete(itemId uint64) {
	i.db.Where("item_id = ?", itemId).Delete(&Item{})
}
func (i *ItemDB) get(itemId uint64) *Item {
	var it Item
	i.db.Where("item_id = ?", itemId).First(&it)
	return &it
}
func (i *ItemDB) getByServiceID(serviceId uint64) []Item {
	var items []Item
	i.db.Where("service_id = ?", serviceId).Find(&items)
	return items
}
func (i *ItemDB) update(it *Item) {
	i.db.Save(it)
}
