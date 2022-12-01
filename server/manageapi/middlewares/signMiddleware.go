package middlewares

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"manageapi/responses"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	timeout      = 300
	signatureKey = "fuck you"
)

// NewSignMiddleware 签名验证中间件
func NewSignMiddleware() web.FilterFunc {
	return func(ctx *context.Context) {

		if !ctx.Input.IsPost() {
			responses.FailMsg(500, ctx)
		}

		timestamp := ctx.Input.Query("timestamp")
		signature := ctx.Input.Query("signature")

		if signature == "" {
			responses.FailMsg(8000, ctx)
		}

		if timestamp == "" {
			responses.FailMsg(8001, ctx)
		}

		timestampInt, err := strconv.ParseInt(timestamp, 10, 64)

		if err != nil {
			responses.FailMsg(8002, ctx)
		}

		dateTime := time.Unix(timestampInt, 0)

		t := time.Now()

		if t.Sub(dateTime).Seconds() > float64(timeout) {
			responses.FailMsg(8003, ctx)
		}

		if signature != Signature(ctx.Request.PostForm) {
			responses.FailMsg(8004, ctx)
		}

		responses.Success(ctx)
	}
}

func Signature(params url.Values) (result string) {

	var b bytes.Buffer
	keys := make([]string, len(params))
	values := make(map[string]string)

	for key, value := range params {
		values[key] = value[0]
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for i, key := range keys {
		if key == "signature" {
			continue
		}

		val := values[key]

		if val != "" && key != "" {
			b.WriteString(key)
			b.WriteString("=")
			b.WriteString(val)
			if i != (len(keys) - 1) {
				b.WriteString("&")
			}
		}
	}

	b.WriteString(signatureKey)
	stringToSign := b.String()

	m := md5.New()
	m.Write([]byte(stringToSign))
	buf := m.Sum(nil)
	result = hex.EncodeToString(buf)

	return result
}
