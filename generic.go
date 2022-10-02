package ginny

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GenericPtr[T interface{}](t *T, errs ...error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		errs = append(errs, nil)
		if t == nil {
			c.AbortIfError(fmt.Errorf("obj is nil: %v, %v, %v, %v, %v, %v", c.Request.URL, c.Request.Method, c.Request.Header, c.Request.Body, c.Request.Form, c.Request.PostForm))
			return
		}
		c.Render(*t, errs[0])
	}
}
func Generic[T interface{}](f func(...string) (T, error), queries ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		var str []string
		for _, query := range queries {
			str = append(str, c.Query(query))
		}
		c.Render(f(str...))
	}
}
