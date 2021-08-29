package model

import (
	"fmt"
	"os"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var DB *xorm.Engine

type MysqlSection struct {
	Addr string `json:"addr"`
	Max  int 	`json:"max"`
	Idle int 	`json:"idle"`
	Debug bool  `json:"debug"`
}

var MySQL = MysqlSection{}

func InitMySQL(MySQL MysqlSection)  {
	conf := MySQL

	db, err := xorm.NewEngine("mysql", conf.Addr)
	if err != nil {
		fmt.Printf("cannot connect mysql[%s]: %v", conf.Addr, err)
		os.Exit(1)
	}

	db.SetMaxIdleConns(conf.Idle)
	db.SetMaxOpenConns(conf.Max)
	db.SetConnMaxLifetime(time.Hour)
	db.ShowSQL(conf.Debug)
	db.Logger().SetLevel(log.LOG_INFO)
	DB = db
}

