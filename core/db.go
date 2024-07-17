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
	dsn := "user:pass@tcp(47.108.94.153:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	peroDB.db = db
	if err != nil {
		peroDB.log.Panic().Err(err).Msg("pero ==> Failed to connect to database")
		panic(err)
	}
	log.Debug().Msg("pero ==> DB start success")
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

type ServiceDB struct {
	db *gorm.DB
}

func newServiceDB() *ServiceDB {
	return &ServiceDB{
		db: getPeroDB().db,
	}
}
func (s *ServiceDB) add(se *Service) {
	s.db.Create(&se)
}
func (s *ServiceDB) update(se *Service) {
	s.db.Save(&se)
}
func (s *ServiceDB) delete(id uint64) {
	var r Service
	s.db.Where("service_id = ?", id).Delete(&r)
}
func (s *ServiceDB) query(id uint64) *Service {
	var r Service
	s.db.Where("service_id = ?", id).Find(&r)
	return &r
}
func (s *ServiceDB) getServiceNum(id uint64) int {
	var r Service
	s.db.Select("num").Find(&r)
	return r.Num
}
func (s *ServiceDB) getAll() *[]Service {
	var t []Service
	s.db.Find(&t)
	return &t
}
