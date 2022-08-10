package main

import (
	_ "adminapi/routers"

	_ "adminapi/inits"
	beego "github.com/beego/beego/v2/server/web"
)

//func dumpHttpFilter(ctx *context.Context) {
//	method := ctx.Request.Method
//	header := ctx.Request.URL
//	body := ctx.Request.Body
//
//	beego.Debug("[dump http] method: ", method, "url: ", header, "body: ", body)
//}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		//beego.InsertFilter("/v1/*", beego.BeforeRouter, dumpHttpFilter)
	}
	beego.Run()
}
