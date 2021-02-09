package e

const (
	Success       = 0
	Error         = 500
	Fail          = 500
	InvalidParams = 400

	NotFoundId = 10001
	FailAuthCheckToken = 10002
	FailAuthCheckTokenTimeout = 10003
	FailAuthTokenCreate = 10004
	FailAuthToken = 10005
	FailLoginStatus = 10006

	// user 相关
	CreateUserFail = 20001
	UserExist      = 20002
	UpdateUserFail = 20004
	UserNameExist  = 20003

	// 地址相关
	CreateAddressFail = 30001
	CreateAddressLimit = 30002
)

var MsgFlags = map[int]string{
	Success:       "success",
	Fail:          "Fail",
	InvalidParams: "请求参数错误",

	NotFoundId: "缺少id",
	FailAuthCheckToken: "鉴权失败",
	FailAuthCheckTokenTimeout: "token 超时",
	FailAuthTokenCreate: "创建 token 失败",
	FailAuthToken: "token 错误",
	FailLoginStatus: "登录状态失效",

	CreateUserFail: "创建账号失败",
	UserExist:      "用户已存在",
	UserNameExist:  "用户名已存在",
	UpdateUserFail: "更新用户失败",

	CreateAddressFail: "创建地址失败",
	CreateAddressLimit: "用户最多可以只能设置 20 个地址",
}
