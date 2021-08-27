package controller

import (
	"github.com/gin-gonic/gin"
	"myblog/service"
)

func LoginGet(c *gin.Context)  {
	//拿到文章的全部数量

	
}

//登录请求
func LoginPost(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag :=service.UserIsExist()
}