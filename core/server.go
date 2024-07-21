package core

import (
	"pero/pkg/http"
)

type Pero struct {
	status bool
	r      *http.PeroHttp
}

var pero Pero

func New() *Pero {
	pero = Pero{
		status: false,
		r:      http.New(),
	}
	return &pero
}

func (p *Pero) GetRoute() *http.PeroHttp {
	return pero.r
}
func (p *Pero) Run() {
	api := NewAPI(pero.r)
	api.Route(pero.r)
	err := api.r.Run()
	if err != nil {
		api.Logger.Error().Err(err).Msg("api run error")
		panic(err)
	}
}
