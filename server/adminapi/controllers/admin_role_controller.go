package controllers

import (
	"adminapi/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRoleController struct {
	beego.Controller
}

func (aR *AdminRoleController) List() {

	res := map[string]interface{}{
		"code":    200,
		"message": "成功",
		"data":    make(map[string]interface{}),
	}
	fmt.Println(res)
	_, adminRoleList, err := models.GetAdminRoleList()

	if err == nil {
		res["data"] = adminRoleList
	}
	aR.Data["json"] = res
	aR.ServeJSON()

}

func (aR *AdminRoleController) Add() {

	adminRoleName := aR.GetString("admin_role_name", "")
	adminPermissionId, _ := aR.GetInt("admin_permission_id", 0)
	storeId, _ := aR.GetInt("store_id", 0)
	id, err := models.AddAdminRole(adminRoleName, adminPermissionId, storeId)
	res := map[string]interface{}{
		"code":    200,
		"message": "成功",
		"data":    make(map[string]interface{}),
	}
	if err == nil && id > 0 {
		aR.Data["json"] = res
	} else {
		res["code"] = 400
		res["message"] = "失败"
	}
	aR.ServeJSON()
}
