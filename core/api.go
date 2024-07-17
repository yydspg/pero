package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"pero/pkg/http"
	"strconv"
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
func (a *API) Route(p *http.PeroHttp) {
	route := p.Group("/v1")
	{
		route.POST("/service/add", a.addService)
		route.GET("/service/get", a.getService)
		route.GET("/service/list", a.listService)
		route.GET("/service/del", a.delService)
		route.POST("/item/add", a.addItem)
		route.GET("/item/list", a.listItems)
		route.POST("/item/del", a.delItem)
		route.POST("/item/add", a.getItem)
		route.POST("/item/update", a.updateItem)
		route.GET("/link/:url", a.link)
		route.POST("/dest/:url", a.dest)
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
	err = a.serviceDB.add(service)
	if valid(err) {
		p.ResponseFail("service add error")
		return
	}
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
	serviceID := p.Param("id")
	ID, err := strconv.ParseUint(serviceID, 10, 64)
	if valid(err) {
		p.ResponseFail("service id invalid")
		return
	}
	num, err := a.serviceDB.getServiceNum(ID)
	if valid(err) {
		p.ResponseFail("query error")
		return
	}
	if num != 0 {
		p.ResponseFail("item nums not null")
		return
	}
	err = a.serviceDB.delete(ID)
	if valid(err) {
		p.ResponseFail("service delete error")
		return
	}
	p.ResponseOK()
}
func (a *API) listService(p *http.Context) {
	t, err := a.serviceDB.getAll()
	if valid(err) {
		p.ResponseFail("service list error")
		return
	}
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
	item := buildItemFromCreate(&req)
	err = a.itemDB.insert(item)
	if valid(err) {
		p.ResponseFail("item add error")
		return
	}
	p.ResponseOK()
}
func (a *API) delItem(p *http.Context) {
	var req struct {
		ServiceID uint64 `json:"service_id"`
		ItemID    uint64 `json:"item_id"`
	}
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json")
		return
	}
	serviceID := req.ServiceID
	itemID := req.ItemID
	if invalid(serviceID) || invalid(itemID) {
		p.ResponseFail("serviceID or itemId is null")
		return
	}
	service := a.serviceDB.query(serviceID)
	if invalid(service) {
		p.ResponseFail("serviceID invalid or build service before")
		return
	}
	err = a.itemDB.delete(itemID)
	if valid(err) {
		p.ResponseFail("item del error")
		return
	}
	p.ResponseOK()
	return
}
func (a *API) updateItem(p *http.Context) {
	var req ItemUpdateReq
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json")
		return
	}
	if req.ServiceID == 0 || req.ItemID == 0 || req.Version == 0 {
		p.ResponseFail("serviceID,DestUrl or version be null")
		return
	}
	service := a.serviceDB.query(req.ServiceID)
	if invalid(service) {
		p.ResponseFail("serviceID invalid or build service before")
		return
	}
	err = a.itemDB.update(buildItemFromUpdate(&req))
	if valid(err) {
		p.ResponseFail("item update error")
		return
	}
	p.ResponseOK()
}
func (a *API) getItem(p *http.Context) {
	var req struct {
		ServiceID uint64 `json:"service_id"`
		ItemID    uint64 `json:"item_id"`
	}
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json")
		return
	}
	serviceID := req.ServiceID
	itemID := req.ItemID
	if invalid(serviceID) || invalid(itemID) {
		p.ResponseFail("serviceID or itemId is null")
		return
	}
	service := a.serviceDB.query(serviceID)
	if invalid(service) {
		p.ResponseFail("serviceID invalid or build service before")
		return
	}
	item, err := a.itemDB.get(itemID)
	if valid(err) {
		a.Log().Err(err)
		p.ResponseFail("item get error")
		return
	}
	p.ResponseOKWithData(item)
}

func (a *API) listItems(p *http.Context) {
	var req struct {
		ServiceID uint64 `json:"service_id"`
	}
	err := p.BindJSON(&req)
	if valid(err) {
		p.ResponseFail("no valid json")
		return
	}
	serviceID := req.ServiceID
	if invalid(serviceID) {
		p.ResponseFail("serviceID invalid or build service before")
		return
	}
	items, err := a.itemDB.getByServiceID(serviceID)
	if valid(err) {
		p.ResponseFail("item list error")
		return
	}
	p.ResponseOKWithData(items)
}
func (a *API) link(p *http.Context) {
	param := p.Param("url")
	if isEmpty(param) {
		p.ResponseFail("destUrl is null")
		return
	}
	link, err := a.itemDB.getLink(param)
	if valid(err) {
		p.ResponseFail("link get error")
		return
	}
	p.ResponseOKWithData(link)
}
func (a *API) dest(p *http.Context) {
	param := p.Param("url")
	if isEmpty(param) {
		p.ResponseFail("shortUrl is null")
		return
	}
	dest, err := a.itemDB.getDest(param)
	if valid(err) {
		p.ResponseFail("dest get error")
		return
	}
	p.ResponseOKWithData(dest)
}
