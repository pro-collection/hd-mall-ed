package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	NotFoundId: "缺少id",

	CreateUserFail: "创建账号失败",
	UserExist:      "用户已存在",
	UserNameExist:  "用户名已存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
