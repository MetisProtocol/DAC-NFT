package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"metisV1.0/conf"
)

func RegisterDatabase() {
	database := beego.AppConfig.String("db_database")
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")
	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci", database)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(createDB)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("1111")
	}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, database)
	orm.RegisterDataBase("default", "mysql", dataSource)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func RegisterModel() {
	fmt.Println("111")
	orm.RegisterModelWithPrefix(conf.GetDatabasePrefix(),
		new(Users),
		new(Dac),
		new(DacAddr),
		new(Eth),
		new(TokenStaking),
	)
}
