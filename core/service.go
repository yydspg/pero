package core

import (
	"time"
)

type Service struct {
	ID          uint64    `db:"id" json:"id"`
	ServiceID   uint64    `db:"service_id" json:"service_id"`
	ServiceName string    `db:"service_name" json:"service_name"`
	Tag         string    `db:"tag" json:"tag"`
	Num         int       `db:"num" json:"num"`
	Status      int8      `db:"status" json:"status"`
	CreateAt    time.Time `db:"create_at" json:"create_at" gorm:"column:create_at;type:timestamp"`
	UpdateAt    time.Time `db:"update_at" json:"update_at" gorm:"column:update_at;type:timestamp"`
}
type ServiceReq struct {
	ServiceName string `db:"service_name" json:"service_name"`
	Tag         string `db:"tag" json:"tag"`
}
type ServiceUpdateReq struct {
	ServiceName string `db:"service_name" json:"service_name"`
	Tag         string `db:"tag" json:"tag"`
	Status      int8   `db:"status" json:"status"`
}

func buildService(r *ServiceReq) *Service {
	var s Service
	s.ServiceName = r.ServiceName
	s.Tag = r.Tag
	s.ServiceID = getID(r.ServiceName)
	s.Status = 0
	s.CreateAt = time.Now()
	s.UpdateAt = time.Now()
	return &s
}
func buildServiceUpdateReq(r *ServiceUpdateReq) *Service {
	var s Service
	s.ServiceName = r.ServiceName
	s.Tag = r.Tag
	s.Status = r.Status
	s.UpdateAt = time.Now()
	return &s
}
