package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/name", api.ChenHua)
	return r
}
