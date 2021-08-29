package service

import (
	"crypto"
	"myblog/util"
)

//判断用户是否存在
func UserIsExist(username,password string)bool  {

	//MD5转换密码
	password = util.MD5(password)

	//查到id
	id :=
}
