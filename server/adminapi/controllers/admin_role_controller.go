package controllers

import (
	"adminapi/models"
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

	_, adminRoleList, err := models.GetAdminRoleList()

	if err != nil {
		res["data"] = adminRoleList
	}
	aR.Data["json"] = res
	aR.ServeJSON()

}
