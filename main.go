package ginny

import (
	"github.com/gin-gonic/gin"
)

// NewRoute 创建路由
// @version 0.1.4
// @description last updated at 9/22/2022 4:58:48 PM
func NewRoute(method, pattern string, handler func(*Context)) Route {
	if method == "" {
		method = "GET"
	}
	return Route{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: ToGinHandler(handler),
	}
}

func NewMiddlewareRoute(method, pattern string, handler ...func(*Context)) MiddlewareRoute {
	if method == "" {
		method = "GET"
	}
	return MiddlewareRoute{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: ToGinHandlers(handler...),
	}
}

func NewGinRoute(method, pattern string, handler gin.HandlerFunc) Route {
	if method == "" {
		method = "GET"
	}
	return Route{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handler,
	}
}

func NewGinMiddlewareRoute(method, pattern string, handler ...gin.HandlerFunc) MiddlewareRoute {
	if method == "" {
		method = "GET"
	}
	return MiddlewareRoute{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handler,
	}
}

func NewContext(c *gin.Context) *Context {
	context := new(Context)
	context.Context = c
	context.SetTimerValue()
	return context
}

// ToGinHandler ext gin.content
func ToGinHandler(h func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(NewContext(c))
	}
}
func ToGinHandlers(h ...func(*Context)) (handlers []gin.HandlerFunc) {
	for _, handler := range h {
		handlers = append(handlers, ToGinHandler(handler))
	}
	return handlers
}
