package core

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PeroDB struct {
	db  *gorm.DB
	log *zerolog.Logger
}

var peroDB *PeroDB

func init() {
	dsn := "user:pass@tcp(47.108.94.153:3306)/pero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	peroDB.db = db
	if err != nil {
		peroDB.log.Panic().Err(err).Msg("pero ==> Failed to connect to database")
		panic(err)
	}
	log.Debug().Msg("pero ==> DB start success")
	pero.status = true
}
func getPeroDB() *PeroDB {
	return peroDB
}

type ItemDB struct {
	db *gorm.DB
}

func newItemDB() *ItemDB {
	return &ItemDB{
		getPeroDB().db,
	}
}
func (i *ItemDB) insert(it *Item) error {
	return i.db.Create(it).Error
}
func (i *ItemDB) delete(itemId uint64) error {
	return i.db.Where("item_id = ?", itemId).Delete(&Item{}).Error
}
func (i *ItemDB) get(itemId uint64) (*Item, error) {
	var it Item
	tx := i.db.Where("item_id = ?", itemId).First(&it)
	return &it, tx.Error
}
func (i *ItemDB) getByServiceID(serviceId uint64) (*[]Item, error) {
	var items []Item
	tx := i.db.Where("service_id = ?", serviceId).Find(&items)
	return &items, tx.Error
}
func (i *ItemDB) update(it *Item) error {
	return i.db.Save(it).Error
}
func (i *ItemDB) getLink(destURL string) (string, error) {
	var item Item
	tx := i.db.Where("dest_url = ?", destURL).Find(&item)
	return item.ShortUrl, tx.Error
}
func (i *ItemDB) getDest(shortURL string) (string, error) {
	var item Item
	tx := i.db.Where("short_url = ?", shortURL).Find(&item)
	return item.ShortUrl, tx.Error
}

type ServiceDB struct {
	db *gorm.DB
}

func newServiceDB() *ServiceDB {
	return &ServiceDB{
		db: getPeroDB().db,
	}
}
func (s *ServiceDB) add(se *Service) error {
	return s.db.Create(&se).Error
}
func (s *ServiceDB) update(se *Service) error {
	return s.db.Save(&se).Error
}
func (s *ServiceDB) delete(id uint64) error {
	var r Service
	return s.db.Where("service_id = ?", id).Delete(&r).Error
}
func (s *ServiceDB) query(id uint64) *Service {
	var r Service
	s.db.Where("service_id = ?", id).Find(&r)
	return &r
}
func (s *ServiceDB) getServiceNum(id uint64) (int, error) {
	var r Service
	tx := s.db.Select("num").Find(&r)
	return r.Num, tx.Error
}
func (s *ServiceDB) getAll() (*[]Service, error) {
	var t []Service
	tx := s.db.Find(&t)
	return &t, tx.Error
}
