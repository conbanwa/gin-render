package ginny

import (
	"strconv"
	"time"
)

//SetMapValue set map value
func (c *Context) SetMapValue(key string) {
	c.Set(mapField, key)
}

//SetTimerValue set timer value
func (c *Context) SetTimerValue() {
	c.Set(timerField, time.Now().UnixNano())
}

// call c.SetKeyValue before call this function
func (c *Context) Render(obj interface{}, err error) {
	field, ok := c.Keys[mapField].(string)
	if !ok {
		if _, ok := obj.([]interface{}); ok {
			field = listField
		} else {
			field = resultField
		}
	}
	r := BaseResult{
		Code: 0,
	}
	if err != nil {
		r.Msg = err.Error()
		r.Code = 1
	}
	if field == noField {
		r.Data = obj
	} else {
		r.Data = map[string]interface{}{field: obj}
	}
	var timer int64
	if start, ok := c.Keys[timerField].(int64); ok {
		timer = time.Now().UnixNano() - start
		r.TimeCost = timer
	}
	c.JSON(200, r)
}

func (c *Context) JsonDirect(rows interface{}, err error) {
	c.SetMapValue(noField)
	c.Render(rows, err)
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
	}
}
