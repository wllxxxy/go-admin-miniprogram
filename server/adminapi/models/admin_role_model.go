package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"time"
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

func AddAdminRole(adminRoleName string, adminPermissionId int, storeId int, operationAdminIp string) (int64, error) {
	adminRoleCreateTime := time.Now()
	o := orm.NewOrm()
	adminRole := AdminRole{
		AdminRoleName:       adminRoleName,
		AdminPermissionId:   adminPermissionId,
		StoreId:             storeId,
		AdminRoleStatus:     1,
		OperationAdminId:    1,
		OperationAdminName:  "系统",
		OperationAdminIp:    operationAdminIp,
		AdminRoleCreateTime: adminRoleCreateTime.Format("2006-01-02 15:04:05"),
	}

	id, err := o.Insert(&adminRole)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}

}

func EditAdminRole(adminRoleId int64, adminRoleName string, adminPermissionId int, storeId int, operationAdminIp string) (int64, error) {
	adminRoleCreateTime := time.Now()
	o := orm.NewOrm()

	adminRoleModel := AdminRole{AdminRoleId: adminRoleId}

	if o.Read(&adminRoleModel) == nil {
		adminRoleModel.AdminRoleName = adminRoleName
		adminRoleModel.AdminPermissionId = adminPermissionId
		adminRoleModel.StoreId = storeId
		adminRoleModel.AdminRoleStatus = 1
		adminRoleModel.OperationAdminId = 1
		adminRoleModel.OperationAdminName = "系统"
		adminRoleModel.OperationAdminIp = operationAdminIp
		adminRoleModel.AdminRoleCreateTime = adminRoleCreateTime.Format("2006-01-02 15:04:05")
		if num, err := o.Update(&adminRoleModel); err == nil {
			return num, nil
		} else {
			return 0, err
		}
	}

	return 0, nil
}

func GetAdminRole(adminRoleId int) (int64, []orm.Params, error) {
	var list []orm.Params

	handler, err := orm.NewQueryBuilder(beego.AppConfig.DefaultString("dbdrivername", ""))

	if err != nil {
		return 0, []orm.Params{}, err
	}

	handler.Select("*").
		From(GetTableName("admin_role")).Where(fmt.Sprintf("admin_role_id = %d", adminRoleId))

	sql := handler.String()

	o := orm.NewOrm()

	num, err := o.Raw(sql).Values(&list)

	if err == nil && num > 0 {
		return num, list, err
	}

	return 0, []orm.Params{}, err
}

func DelAdminRole(adminRoleId int64) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Delete(&AdminRole{AdminRoleId: adminRoleId})

	if err == nil && num > 0 {
		return num, err
	}

	return 0, nil
}
