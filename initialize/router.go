package initialize

import (
	"bubble/controller"
	"bubble/routers"
	"bubble/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	PublicGroup := r.Group("")
	// PrivateGroup := r.Group("")
	{
		routers.V1Router(PublicGroup)    // v1基础路由
		routers.V2Router(PublicGroup)    // v1基础路由
		routers.StudyRouter(PublicGroup) //study基础路由
	}

	return r

}
