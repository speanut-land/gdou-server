package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/speanut-land/gdou-server/docs"
	"github.com/speanut-land/gdou-server/routers/api"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/register", api.Register)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
