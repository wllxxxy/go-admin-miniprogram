package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRole struct {
	AdminRoleId         int64 `orm:"pk;column(role_id)"`
	AdminRoleName       string
	AdminPermissionId   int
	StoreId             int
	AdminRoleStatus     int
	OperationAdminId    int
	OperationAdminName  string
	OperationAdminIp    string
	AdminRoleCreateTime string
}

func (aR *AdminRole) TableName() string {
	return GetTableName("admin_role")
}

func init() {
	orm.RegisterModel(new(AdminRole))
}

func GetAdminRoleList() (int64, []orm.Params, error) {
	var list []orm.Params

	handler, err := orm.NewQueryBuilder(beego.AppConfig.DefaultString("dbdrivername", ""))

	if err != nil {
		return 0, []orm.Params{}, err
	}

	handler.Select("*").
		From(GetTableName("admin_role"))

	sql := handler.String()

	o := orm.NewOrm()

	num, err := o.Raw(sql).Values(&list)

	if err == nil && num > 0 {
		return num, list, err
	}

	return 0, []orm.Params{}, err
}
