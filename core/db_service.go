package core

import (
	"gorm.io/gorm"
)

type ServiceDB struct {
	db *gorm.DB
}

func newServiceDB() *ServiceDB {
	return &ServiceDB{
		db: getPeroDB().db,
	}
}
func (s *ServiceDB) add(se *Service) {
	s.db.Create(se)
}
func (s *ServiceDB) update(se *Service) {
	s.db.Save(se)
}
func (s *ServiceDB) delete(se *Service) {
	s.db.Delete(se)
}
func (s *ServiceDB) query(se *Service) *gorm.DB {
	return s.db.Where(se)
}
