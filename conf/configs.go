package conf

import "github.com/astaxie/beego"

func GetDatabasePrefix() string {
	return beego.AppConfig.DefaultString("db_prefix", "metis_")
}
