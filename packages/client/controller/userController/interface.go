package userController

type authStruct struct {
	Username string `json:"username" valid:"required~缺少用户名"`
	Password string `json:"password" valid:"required~缺少用户密码"`
}
