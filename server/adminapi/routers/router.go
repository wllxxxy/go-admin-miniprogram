// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"adminapi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 后台用户
	beego.Router("/admin/list", &controllers.AdminController{}, "get:List")
	//beego.Router("/admin/detail?:admin_id", &controllers.AdminController{}, "get:Detail")
	//beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
	//beego.Router("/admin/delete?:admin_id", &controllers.AdminController{}, "get:Delete")
	//beego.Router("/admin/edit?:admin_id", &controllers.AdminController{}, "get:Edit")

	// 后台角色
	beego.Router("/admin_role/list", &controllers.AdminRoleController{}, "get:List")
	//beego.Router("/admin_role/detail?:role_id", &controllers.AdminController{}, "get:Detail")
	//beego.Router("/admin_role/add", &controllers.AdminController{}, "get:Add")
	//beego.Router("/admin_role/delete?:role_id", &controllers.AdminController{}, "get:Delete")
	//beego.Router("/admin_role/edit?:role_id", &controllers.AdminController{}, "get:Edit")
}
