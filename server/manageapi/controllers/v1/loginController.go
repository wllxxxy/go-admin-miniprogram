package v1

type LoginController struct {
	BaseController
}

// Login 登陆
func (l *LoginController) Login() {
	l.Data["json"] = map[string]interface{}{}
	l.ServeJSON()
}
