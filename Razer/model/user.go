package model

import (
	"fmt"
)

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	portrait string `json:"portrait"`
	Status   string `json:"status"`
	CreateAt int 	`json:"create_at"`
}

func UserGet(where string,args ...interface{})(*User,error) {
	var obj User
	has, err := DB.Where(where, args...).Get(&obj)
	if err != nil{
		fmt.Print("unable to get user")
	}

	if !has{
		return nil,nil
	}

	return &obj,nil
}

func UserGetByUsername(username string)(*User,error)  {
	return UserGet("username = ?",username)
}