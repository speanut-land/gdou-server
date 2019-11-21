package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/speanut-land/gdou-server/docs"
	"github.com/speanut-land/gdou-server/routers/api"
	"github.com/speanut-land/gdou-server/routers/api/sendcode"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := r.Group("/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
		user.POST("/resetPassword", api.ResetPassword)
	}

	sendCode := r.Group("/sendCode")
	{
		sendCode.POST("/register", sendcode.Register)
		sendCode.POST("/resetPassword", sendcode.ResetPassword)
	}
	return r
}
