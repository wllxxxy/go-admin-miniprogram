package responses

import (
	"github.com/beego/beego/v2/server/web/context"
	"manageapi/utils"
	"net/http"
)

type ResponseVo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	FAIL = iota
	SUCCESS
)

func Result(code int, data interface{}, msg string, ctx *context.Context) {
	ctx.Output.SetStatus(http.StatusOK)
	ctx.Output.JSON(ResponseVo{
		code,
		msg,
		data,
	}, false, false)
}

func Success(ctx *context.Context) {
	Result(SUCCESS, map[string]interface{}{}, "成功", ctx)
}

func SuccessData(data interface{}, ctx *context.Context) {
	Result(SUCCESS, data, "成功", ctx)
}

func Fail(ctx *context.Context) {
	Result(FAIL, map[string]interface{}{}, "失败", ctx)
}

func FailMsg(code int, ctx *context.Context) {
	Result(code, map[string]interface{}{}, utils.GetErrorCodeMessage(code), ctx)
}
