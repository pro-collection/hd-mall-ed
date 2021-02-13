package userController

type updateParamsStruct struct {
	Name     string `json:"name"  valid:"required~缺少用户名"`
	Password string `json:"password" valid:"required~缺少用户密码"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Role     string `json:"role" valid:"required~缺少角色"`
}
