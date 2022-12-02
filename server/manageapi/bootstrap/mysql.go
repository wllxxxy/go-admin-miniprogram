package bootstrap

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"manageapi/conf"
)

func init() {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)

	if err != nil {
		logs.Error("mysql register driver error:", err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.CONFIG.Mysql.Username,
		conf.CONFIG.Mysql.Password,
		conf.CONFIG.Mysql.Host,
		conf.CONFIG.Mysql.Port,
		conf.CONFIG.Mysql.Database,
	)

	runmode, _ := config.String("runmode")

	if runmode == "dev" {
		orm.Debug = true
	}

	err = orm.RegisterDataBase("default", "mysql", dataSource)

	if err != nil {
		logs.Error("mysql register database error:", err)
	}
}
