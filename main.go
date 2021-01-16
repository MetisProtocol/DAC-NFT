package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"metisV1.0/models"
	_ "metisV1.0/routers"
)

func init() {
	models.RegisterDatabase()
	models.RegisterModel()
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
