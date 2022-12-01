package v1

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ExampleController struct {
	beego.Controller
}

func (e *ExampleController) Index() {
	e.Data["json"] = map[string]interface{}{}
	e.ServeJSON()
}
