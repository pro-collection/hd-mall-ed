package authController

type logParamsStruct struct {
	Name     string `json:"name" valid:"required~缺少用户名"`
	Password string `json:"password" valid:"required~缺少用户密码"`
}

var adminTokenKey = "admin-token"
