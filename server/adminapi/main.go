package main

import (
	_ "adminapi/inits"
	_ "adminapi/routers"
	beego "github.com/beego/beego/v2/server/web"
)
//测试文件
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
