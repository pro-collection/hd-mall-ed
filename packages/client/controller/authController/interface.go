package authController

type authStruct struct {
	Name     string `json:"name" valid:"required~缺少用户名"`
	Password string `json:"password" valid:"required~缺少用户密码"`
}
