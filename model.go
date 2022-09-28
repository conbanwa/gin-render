package ginny

import "github.com/gin-gonic/gin"

const (
	mapField     = "ctx"
	noField      = "result"
	resultField  = "result"
	listField    = "list"
	methodField  = "method"
	patternField = "pattern"
	timerField   = "timer"
)

type (
	// Context ext content struct
	Context struct {
		*gin.Context
	}

	BaseResult struct {
		Code     int         `json:"code,omitempty"`
		Msg      string      `json:"msg,omitempty"`
		Data     interface{} `json:"data,omitempty"`
		TimeCost int64       `json:"timeCost,omitempty"`
		Total    int64       `json:"total,omitempty"`
		// Key    string            `json:"key,omitempty"`
		// Detail map[string]string `json:"detail,omitempty"`
	}
	// Route is the information for every URI.
	Route struct {
		// Method is the string for the HTTP method. ex) GET, POST etc..
		Method string
		// Pattern is the pattern of the URI.
		Pattern string
		// HandlerFunc is the handler function of this route.
		HandlerFunc gin.HandlerFunc
	}
	// MiddlewareRoute is the information for every URI.
	MiddlewareRoute struct {
		// Method is the string for the HTTP method. ex) GET, POST etc..
		Method string
		// Pattern is the pattern of the URI.
		Pattern string
		// HandlerFunc is the handler function of this route.
		HandlerFunc []gin.HandlerFunc
	}
)
