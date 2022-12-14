package service

import (
	"github.com/go-xorm/xorm"
	"log"
	"fmt"
	"errors"
	"../model"
)

var DbEngin *xorm.Engine
func  init()  {
	drivename :="mysql"
	DsName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin,err = xorm.NewEngine(drivename,DsName)
	if nil!=err && ""!=err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动User
	DbEngin.Sync2(new(model.User))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}

