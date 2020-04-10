package helpers

import (
	"github.com/gin-gonic/gin"
)

//Response is the model of http response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Respond is the default function used to create responses
func Respond(c *gin.Context, code int, message string, data interface{}) {

	if data != nil {
		c.JSON(code, map[string]interface{}{
			"code": code,
			"msg":  message,
			"data": data,
		})
		return
	}

	c.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  message,
	})
}
