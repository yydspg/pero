package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"pero/pkg/http"
	"sync/atomic"
)

type API struct {
	sID atomic.Uint64
	tID atomic.Uint64
	zerolog.Logger
	r         *gin.Engine
	serviceDB *ServiceDB
	itemDB    *ItemDB
}

func NewAPI() *API {
	return &API{
		Logger:    zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}),
		r:         gin.New(),
		serviceDB: newServiceDB(),
		itemDB:    newItemDB(),
	}
}
func Route(p *http.PeroHttp) {
	route := p.Group("/link")
	{
		route.POST("/service/add")
		route.GET("/service/get")
		route.GET("/service/list")
		route.GET("/service/del")
		route.POST("/service/update")
		route.GET("/service/items")
		route.POST("/service/del")
		route.POST("/get")
	}
}
func (a *API) addService(p *http.Context) {
	var req ServiceReq
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json ")
		return
	}
	if isEmpty(req.ServiceName) || isEmpty(req.Tag) {
		p.ResponseFail("service name or tag is empty")
	}
	service := buildService(&req)
	a.serviceDB.add(service)
	p.ResponseOK()
}
func (a *API) getService(p *http.Context) {
	serviceID := p.GetUint64("service_id")
	s := a.serviceDB.query(serviceID)
	if invalid(s) {
		p.ResponseFail("no service info")
		return
	}
	p.ResponseData(200, s)
}
func (a *API) delService(p *http.Context) {
	serviceID := p.GetUint64("service_id")
	num := a.serviceDB.getServiceNum(serviceID)
	if num != 0 {
		p.ResponseFail("item nums not null")
		return
	}
	a.serviceDB.delete(serviceID)
	p.ResponseOK()
}
func (a *API) listService(p *http.Context) {
	t := a.serviceDB.getAll()
	p.ResponseOKWithData(*t)
}
func (a *API) addItem(p *http.Context) {
	var req ItemReq
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json")
		return
	}
	if req.ServiceID == 0 || isEmpty(req.DestUrl) || req.Version == 0 {
		p.ResponseFail("serviceID,DestUrl or version be null")
		return
	}
	service := a.serviceDB.query(req.ServiceID)

	if invalid(service) {
		p.ResponseFail("serviceID invalid or build service before")
		return
	}
	service.Num++
	item := buildItem(&req)

}
