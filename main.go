package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/speanut-land/gdou-server/models"
	"github.com/speanut-land/gdou-server/pkg/logging"
	"github.com/speanut-land/gdou-server/pkg/redis"
	"github.com/speanut-land/gdou-server/pkg/setting"
	"github.com/speanut-land/gdou-server/pkg/util"
	"github.com/speanut-land/gdou-server/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	_ = redis.SetUp()
	logging.Setup()
	util.Setup()
}
func main() {
	//设置gin框架的运行模式
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] 开始启动http服务 端口号为 %s", endPoint)

	server.ListenAndServe()
}
