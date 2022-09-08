package controllers

import (
	"adminapi/models"
	"common/libs"
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
	operationAdminIp := libs.RemoteIp(aR.Ctx.Request)
	id, err := models.AddAdminRole(adminRoleName, adminPermissionId, storeId, operationAdminIp)
	res := map[string]interface{}{
		"code":    200,
		"message": "添加成功",
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

func (aR *AdminRoleController) Edit() {
	adminRoleName := aR.GetString("admin_role_name", "")
	adminPermissionId, _ := aR.GetInt("admin_permission_id", 0)
	storeId, _ := aR.GetInt("store_id", 0)
	adminRoleId, _ := aR.GetInt64("admin_role_id", 0)
	operationAdminIp := libs.RemoteIp(aR.Ctx.Request)
	num, err := models.EditAdminRole(adminRoleId,adminRoleName, adminPermissionId, storeId, operationAdminIp)
	res := map[string]interface{}{
		"code":    200,
		"message": "修改成功",
		"data":    make(map[string]interface{}),
	}
	if err == nil && num > 0 {
		aR.Data["json"] = res
	} else {
		res["code"] = 400
		res["message"] = "修改失败"
	}
	aR.ServeJSON()
}

func (aR *AdminRoleController) Detail() {
	adminRoleId, _ := aR.GetInt("admin_role_id", 0)
	if adminRoleId != 0 {
		_, adminRole, err := models.GetAdminRole(adminRoleId)

		res := map[string]interface{}{
			"code":    200,
			"message": "成功",
			"data":    make(map[string]interface{}),
		}
		if err == nil {
			res["data"] = adminRole[0]
		}
		aR.Data["json"] = res
		aR.ServeJSON()
	}

}


func (aR *AdminRoleController) Delete() {
	adminRoleId, _ := aR.GetInt64("admin_role_id", 0)
	if adminRoleId != 0 {
		num, err := models.DelAdminRole(adminRoleId)

		res := map[string]interface{}{
			"code":    200,
			"message": "删除成功",
			"data":    make(map[string]interface{}),
		}
		if err == nil && num > 0 {
			aR.Data["json"] = res
		} else {
			res["code"] = 400
			res["message"] = "修改失败"
		}
		aR.ServeJSON()
	}

}
