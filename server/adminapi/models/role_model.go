package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type Role struct {
	RoleId       int64 `orm:"pk;column(role_id)"`
	RoleName     string
	RoleNameCode string
	RoleStatus   int
	RoleMark     string
	AddTime      int
	AddUserId    int
	EditTime     int
	EditUserId   int
}

func (r *Role) TableName() string {
	return GetTableName("role")
}

func init() {
	orm.RegisterModel(new(Role))
}

func List() (int64, []orm.Params, error) {
	var list []orm.Params

	handler, err := orm.NewQueryBuilder(beego.AppConfig.DefaultString("dbdrivername", ""))

	if err != nil {
		return 0, []orm.Params{}, err
	}

	handler.Select("*").
		From(GetTableName("role"))

	sql := handler.String()

	o := orm.NewOrm()

	num, err := o.Raw(sql).Values(&list)

	if err == nil && num > 0 {
		return num, list, err
	}

	return 0, []orm.Params{}, err
}
