package core

type Service struct {
	ID          int64  `db:"id" json:"id"`
	ServiceID   int64  `db:"service_id" json:"service_id"`
	ServiceName string `db:"service_name" json:"service_name"`
	Tag         string `db:"tag" json:"tag"`
	Status      int8   `db:"status" json:"status"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
}
