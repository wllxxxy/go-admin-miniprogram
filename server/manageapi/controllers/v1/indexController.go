package v1

import (
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

// Index 首页
func (i *IndexController) Index() {
	i.Data["json"] = map[string]interface{}{}
	i.ServeJSON()
}

// Logout 退出
func (i *IndexController) Logout() {
	i.Data["json"] = map[string]interface{}{}
	i.ServeJSON()
}
