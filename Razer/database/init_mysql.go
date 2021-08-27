package database

import "database/sql"

//DB 数据库
var DB *sql.DB

//InitMysql 初始化mysql数据库
func InitMysql()  {
	//连接数据库
	if DB == nil{
		DB, _ =sql.Open()

	}

}
