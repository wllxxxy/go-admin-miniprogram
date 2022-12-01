package utils

func GetErrorCodeMessage(code int) string {

	//var msg string

	message := map[int]string{
		-1:   "未知",
		0:    "失败",
		1:    "成功",
		500:  "非法请求",
		8000: "signature参数丢失",
		8001: "timestamp参数丢失",
		8002: "timestamp格式错误",
		8003: "url失效",
		8004: "signature验证失败",
	}

	msg, ok := message[code]

	if ok == false {
		msg = message[-1]
	}

	return msg
}
