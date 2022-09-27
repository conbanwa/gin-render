package ginny

import "github.com/gin-gonic/gin"

func Generic[T interface{}](obj T, errs ...error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		errs = append(errs, nil)
		c.Render(obj, errs[0])
	}
}

func GenericQuery[T interface{}](f func(...string) (T, error), queries ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		var str []string
		for _, query := range queries {
			str = append(str, c.Query(query))
		}
		c.Render(f(str...))
	}
}
