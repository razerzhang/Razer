package service

import (
	"myblog/model"
	"myblog/util"
)

//判断用户是否存在
func UserIsExist(username,password string)bool  {

	//MD5转换密码
	password = util.MD5(password)

	//通过username查询
	user,err := model.UserGetByUsername(username)
	if err != nil{
		return false
	}
	if user != nil {
		return true
	}else {
		return false
	}
}
