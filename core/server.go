package core

import (
	"pero/pkg/http"
	"sync/atomic"
)

type Pero struct {
	sID    atomic.Uint64
	iID    atomic.Uint64
	status bool
	r      *http.PeroHttp
}

var pero Pero

func New() *Pero {
	pero = Pero{
		sID:    atomic.Uint64{},
		iID:    atomic.Uint64{},
		status: false,
		r:      http.New(),
	}
	pero.sID.Store(0)
	pero.iID.Store(0)
	return &pero
}

func (p *Pero) GetRoute() *http.PeroHttp {
	return pero.r
}
