package middlewares

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"manageapi/conf"
	"manageapi/utils"
	"strings"
)

// NewAuthenticateMiddleware 登陆验证中间件
func NewAuthenticateMiddleware() web.FilterFunc {
	return func(ctx *context.Context) {
		urlAllow := conf.CONFIG.Common.UrlAllow
		url := strings.ToLower(strings.TrimPrefix(ctx.Input.URL(), "/v1/"))

		// JWT令牌验证
		if utils.InSlice(url, urlAllow) == false {

		}
	}
}
