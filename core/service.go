package core

type Service struct {
	ID          uint64 `db:"id" json:"id"`
	ServiceID   uint64 `db:"service_id" json:"service_id"`
	ServiceName string `db:"service_name" json:"service_name"`
	Tag         string `db:"tag" json:"tag"`
	Num         int    `db:"num" json:"num"`
	Status      int8   `db:"status" json:"status"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
}
type ServiceReq struct {
	ServiceName string `db:"service_name" json:"service_name"`
	Tag         string `db:"tag" json:"tag"`
}

func buildService(r *ServiceReq) *Service {
	var s Service
	s.ServiceName = r.ServiceName
	s.Tag = r.Tag
	s.ServiceID = pero.sID.Add(1)
	s.Status = 0
	return &s
}
