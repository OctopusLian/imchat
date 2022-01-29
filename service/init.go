/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-28 23:48:54
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-29 22:15:44
 */
package service

import (
	"errors"
	"fmt"
	"imchat/model"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:mysql123@(127.0.0.1:3306)/imchat?charset=utf8mb4"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(false)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动同步表
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
