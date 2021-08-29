package model

import "xorm.io/xorm"

var DB *xorm.Engine

type MysqlSection struct {
	Addr string `json:"addr"`
	Max  int 	`json:"max"`
	Idle int 	`json:"idle"`
	Debug bool  `json:"debug"`
}
func ()  {
	
}
