package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}


// 后台管理用户列表
func (a *AdminController) List() {
	//o := orm.NewOrm()
	//var admin []models.Admin
	//_, err := o.QueryTable("admin").OrderBy("-admin_id").All(&admin)
	//
	//var serve = make(map[string]interface{})
	//
	//if err != nil {
	//	serve["code"] = "1"
	//	serve["message"] = "查询失败"
	//} else {
	//	serve["code"] = "0"
	//	serve["message"] = "查询成功"
	//	serve["data"] = admin
	//}
	//a.Data["json"] = &serve
	//defer a.ServeJSON()
	//adminList := models.GetAdminList()
	//a.Data["json"] = adminList
	//a.ServeJSON()
}


func (a *AdminController) Detail() {
	//adminId := a.GetString(":admin_id")
	//if adminId != "" {
	//	admin, err := models.GetAdmin(adminId)
	//	var serve = make(map[string]interface{})
	//	if err != nil {
	//		serve["code"] = "1"
	//		serve["message"] = "查询失败"
	//	} else {
	//		serve["code"] = "0"
	//		serve["message"] = "查询成功"
	//		serve["data"] = admin
	//	}
	//
	//	a.Data["json"] = &serve
	//	defer a.ServeJSON()
	//}



}
