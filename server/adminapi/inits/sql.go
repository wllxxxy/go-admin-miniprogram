package inits

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user, _ := beego.AppConfig.String("mysqluser")
	passwd, _ := beego.AppConfig.String("mysqlpass")
	host, _ := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname, _ := beego.AppConfig.String("mysqldb")

	if err != nil {
		port = 3306
	}

	runmode, _ := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		orm.Debug = true
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
}
