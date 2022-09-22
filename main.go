package render

import (
	"time"
)

// @version 0.0.2
// @description last updated at 9/22/2022 4:31:33 PM

//SetMapValue set map value
func (c *Context) SetMapValue(key string) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[mapField] = key
}

//SetTimerValue set timer value
func (c *Context) SetTimerValue() {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[timerField] = time.Now().UnixNano()
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
