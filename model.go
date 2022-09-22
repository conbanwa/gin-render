package ginny

const (
	mapField    = "ctx"
	noField     = "result"
	resultField = "result"
	listField   = "list"
	methodField = "method"
	pathField   = "path"
	timerField  = "timer"
)

// Context ext content struct
type (
	// HandlerFunc achieves gin.HandlerFunc
	HandlerFunc func(*Context) string
	BaseResult  struct {
		Code     int         `json:"code,omitempty"`
		Msg      string      `json:"msg,omitempty"`
		Data     interface{} `json:"data,omitempty"`
		TimeCost int64       `json:"timeCost,omitempty"`
		Total    int64       `json:"total,omitempty"`
		// Key    string            `json:"key,omitempty"`
		// Detail map[string]string `json:"detail,omitempty"`
	}
)
