package controller

import (
	"github.com/gin-gonic/gin"
	"myblog/service"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func LoginGet(c *gin.Context)  {
	//拿到文章的全部数量

	
}

//登录请求
func LoginPost(c *gin.Context)  {
	//表单接收数据
	username := c.PostForm("username")
	password := c.PostForm("password")

	//判断用户是否存在
	flag :=service.UserIsExist(username,password)
}