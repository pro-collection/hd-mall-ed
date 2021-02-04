package e

const (
	Success       = 0
	Error         = 500
	Fail          = 500
	InvalidParams = 400

	NotFoundId = 10001

	// user 相关
	CreateUserFail = 20001
	UserExist      = 20002
	UserNameExist  = 20003
	UpdateUserFail = 20004
)

var MsgFlags = map[int]string{
	Success:       "success",
	Fail:          "Fail",
	InvalidParams: "请求参数错误",

	NotFoundId: "缺少id",

	CreateUserFail: "创建账号失败",
	UserExist:      "用户已存在",
	UserNameExist:  "用户名已存在",
	UpdateUserFail: "更新用户失败",
}
