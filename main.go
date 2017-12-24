package main

import (
	_"MyJannah-API/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "myjannah:@dbmyjannah2018@tcp(127.0.0.1:3306)/dbjannah")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

