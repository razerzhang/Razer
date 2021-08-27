package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


// HandlerFunc 作为中间件的返回参数
func UseSession()gin.HandlerFunc  {
	store:=cookie.NewStore([]byte("loginUser"))
	return sessions.Sessions("loginUser", store)
}