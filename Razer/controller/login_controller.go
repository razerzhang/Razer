package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myblog/service"
	"net/http"
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

	//返回前端结果
	resp := gin.H{}
	if flag{
		//创建session
		s := sessions.Default(c)
		//设置session
		s.Set("loginUser",username)
		s.Save()
		resp = gin.H{
			"code":	0,
			"message":"登录成功",
		}
	}else {
		resp = gin.H{
			"code": 1,
			"message":"用户名或密码错误",
		}
	}
	c.JSON(http.StatusOK,resp)
}