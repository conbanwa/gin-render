package ginny

import (
	"github.com/gin-gonic/gin"
)

// @version 0.1.1
// @description last updated at 9/22/2022 4:58:48 PM
func NewRoute(method, pattern string, handler func(*Context)) Route {
	return Route{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: ToGinHandler(handler),
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

// Cors 直接放行所有跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 处理请求
		c.Next()
	}
}
