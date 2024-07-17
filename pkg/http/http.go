package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"sync"
)

type PeroHttp struct {
	r    *gin.Engine
	pool sync.Pool
}
type Context struct {
	*gin.Context
	lg zerolog.Logger
}
type HandlerFunc func(c *Context)
type RouterGroup struct {
	*gin.RouterGroup
	L *PeroHttp
}

func (l *PeroHttp) Group(relativePath string) *RouterGroup {
	return newRouterGroup(l.r.Group(relativePath), l)
}
func newRouterGroup(g *gin.RouterGroup, l *PeroHttp) *RouterGroup {
	return &RouterGroup{RouterGroup: g, L: l}
}
func (l *PeroHttp) handlersToGinHandleFuncs(handlers []HandlerFunc) []gin.HandlerFunc {
	newHandlers := make([]gin.HandlerFunc, 0, len(handlers))
	if handlers != nil {
		for _, handler := range handlers {
			newHandlers = append(newHandlers, l.HttpHandler(handler))
		}
	}
	return newHandlers
}
func (l *PeroHttp) HttpHandler(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		hc := l.pool.Get().(*Context)
		hc.reset()
		hc.Context = c
		defer l.pool.Put(hc)

		handlerFunc(hc)

		//handlerFunc(&Context{Context: c})
	}
}
func (c *Context) reset() {
	c.Context = nil
}
func (r *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.POST(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// GET method
func (r *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	r.RouterGroup.GET(relativePath, r.L.handlersToGinHandleFuncs(handlers)...)
}

// ResponseError ResponseError
func (c *Context) ResponseError(err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    err.Error(),
		"status": http.StatusBadRequest,
	})
}
func (c *Context) ResponseFail(msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":    msg,
		"status": http.StatusBadRequest,
	})
}

// ResponseOK 返回正确
func (c *Context) ResponseOK() {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

// ResponseOKWithData 返回正确并并携带数据
func (c *Context) ResponseOKWithData(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

// ResponseData 返回状态和数据
func (c *Context) ResponseData(status int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   data,
	})
}

// ResponseStatus 返回状态
func (c *Context) ResponseStatus(status int) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

/*
这段代码是为 Gin Web 框架设计的一个封装层，主要目的是为了引入对象池（通过 sync.Pool）以及提供一个自定义的上下文（Context）和处理器函数（HandlerFunc）。以下是各个部分的详细解释：

PeroHttp 结构体
PeroHttp 结构体包含两个字段：

r: 一个指向 gin.Engine 的指针，这是 Gin 框架的核心组件，用于处理 HTTP 请求。
pool: 一个 sync.Pool 实例，用于管理并重复使用 Context 对象，减少内存分配和垃圾收集的开销。
Context 结构体
Context 结构体继承自 Gin 的原生 gin.Context，并添加了一个 lg 字段，用于存储日志记录器。这使得开发者可以在处理器函数中方便地使用日志功能。

HandlerFunc 类型
HandlerFunc 是一个类型别名，表示一个接受 *Context 参数的函数。这是 Gin 的 gin.HandlerFunc 的定制版本，允许开发者使用自定义的 Context 而不是原生的 gin.Context。

RouterGroup 结构体
RouterGroup 结构体包装了 Gin 的 gin.RouterGroup，并添加了一个 L 字段，指向 PeroHttp 实例。这使得路由组可以访问 PeroHttp 的方法和属性。

方法
handlersToGinHandleFuncs 方法将一系列 HandlerFunc 转换为 gin.HandlerFunc 列表，以便 Gin 可以理解并执行这些处理器。
WKHttpHandler 方法是一个辅助函数，用于转换处理器函数。它从 sync.Pool 获取一个 Context 实例，重置它，然后将其与当前的 Gin 上下文关联，并在处理器执行后将其放回池中。
reset 方法用于重置 Context，确保每个请求之间没有状态残留。
POST, GET, DELETE, 和 PUT 方法分别用于注册处理特定 HTTP 方法的路由。它们使用 handlersToGinHandleFuncs 方法转换处理器函数列表，然后将转换后的函数传递给 Gin 的相应方法。
*/
