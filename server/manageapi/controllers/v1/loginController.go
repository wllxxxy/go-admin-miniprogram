package v1

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// Login 登陆
func (l *LoginController) Login() {
	l.Data["json"] = map[string]interface{}{}
	l.ServeJSON()
}
