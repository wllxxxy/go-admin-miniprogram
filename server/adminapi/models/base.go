package models

import beego "github.com/beego/beego/v2/server/web"

func GetTableName(name string) string {
	tableName, _ := beego.AppConfig.String("mysqlpre")
	return tableName + name
}
