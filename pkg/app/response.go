package app

import (
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, success bool, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Msg:     e.GetMsg(errCode),
		Success: success,
		Data:    data,
	})
	return
}
