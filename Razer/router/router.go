package router

import (
	"github.com/gin-gonic/gin"
	"myblog/controller"
)

//路由，即方法返回的就是路由
func InitRoute()*gin.Engine  {
	//创建一个gin实例
	r := gin.Default()

	//静态资源加载
	r.Static("/static","./static")
	r.Use(gin.Logger())

	//登录
	r.GET("/login",controller.LoginGet())
	r.POST("/login",controller.LoginPost())
}

