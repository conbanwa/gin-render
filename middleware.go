package render

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func BatchGet(routerGroup *gin.RouterGroup, handlers ...HandlerFunc) {
	for _, h := range handlers {
		c := new(Context)
		c.Context = new(gin.Context)
		routerGroup.GET(h(nil), h.ToGin())
	}
}

func NewContext(c *gin.Context) *Context {
	context := new(Context)
	context.Context = c
	context.SetTimerValue()
	return context
}

// Handler ext gin.content
func Handler(h func(*Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(NewContext(c))
	}
}

// var store = persistence.NewInMemoryStore(60 * time.Second)

func (h HandlerFunc) ToGin() gin.HandlerFunc {
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

func (c *Context) AbortIfError(err error) bool {
	if err != nil {
		r := BaseResult{Msg: err.Error(), Code: 1}
		c.AbortWithStatusJSON(400, r)
		return true
	}
	return false
}

func (c *Context) QueryInt(str string) int {
	i, err := strconv.Atoi(c.Query(str))
	if err != nil {
		i = 0
	}
	return i
}
func (c *Context) PostInt(str string) int {
	i, err := strconv.Atoi(c.PostForm(str))
	if err != nil {
		i = 0
	}
	return i
}
func RejectAuth() (int, BaseResult) {
	return 403, BaseResult{
		Msg:  "forbidden",
		Code: 1,
		// ErrMsg: ErrForbidden,
	}
}
