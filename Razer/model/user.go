package model

import (
	"github.com/gin-gonic/gin"
	"myblog/database"
)

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	portrait string `json:"portrait"`
	Status   string `json:"status"`
	CreateAt int 	`json:"create_at"`
}

func UserGet(where string,args ...interface{})  {
	var obj User
 	has, err :=
