package ginny

import "github.com/gin-gonic/gin"

func GenericPtr[T interface{}](t *T, errs ...error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		errs = append(errs, nil)
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
