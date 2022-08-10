package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"go-admin-miniprogram/server/common/libs"
	"net/http"
)

type AdminRole struct {
	AdminRoleId         int64 `orm:"pk;column(admin_role_id)"`
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

func AddAdminRole(adminRoleName string, adminPermissionId int, storeId int) (int64, error) {
	var r *http.Request
	ip := RemoteIp(r)
	//ip := libs.ClientPublicIP(r)
	//if ip == ""{
	//	ip = exnet.ClientIP(r)
	//}
	o := orm.NewOrm()
	adminRole := AdminRole{
		AdminRoleName:       adminRoleName,
		AdminPermissionId:   adminPermissionId,
		StoreId:             storeId,
		AdminRoleStatus:     1,
		OperationAdminId:    1,
		OperationAdminName:  "系统",
		OperationAdminIp:    "127.0.0.1",
		AdminRoleCreateTime: "2022-08-10 17:00:00",
	}

	id, err := o.Insert(&adminRole)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}

}
