package ginny

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

//TimeoutMiddleware wrap the restful context with a timeout
func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				// write response and abort the restful
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.JSON(http.StatusOK, BaseResult{
					Code: 2,
					Msg:  "restful TIMEOUT",
				})
				c.Abort()
			}
			//cancel to clear resources after finished
			cancel()
		}()
		// replace restful with context wrapped restful
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
