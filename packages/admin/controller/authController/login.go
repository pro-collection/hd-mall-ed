package authController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/adminAuthModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
)

func Login(c *gin.Context) {
	api := adminApp.ApiInit(c)
	var err error
	params := logParamsStruct{}

	_ = c.ShouldBindJSON(&params)

	// 参数校验
	if api.ValidateHasError(params) {
		return
	}

	data := make(map[string]interface{})
	code := e.FailAuthToken
	id, err := adminAuthModel.CheckAuth(params.Name, params.Password)
	if err != nil {
		api.ResFailMessage(e.FailAuthCheckToken, err.Error())
		return
	}

	if id > 0 {
		token, claims, err := jwtUtil.AdminGenerateToken(params.Name, params.Password, id)

		if err != nil {
			code = e.FailAuthTokenCreate
		} else {
			// 设置 cookie
			api.SetCookie(adminTokenKey, token, int(claims.ExpiresAt))
			code = e.Success
		}
	}

	if code == e.Success {
		api.Response(data)
		return
	}

	// error 的返回
	api.ResFail(code)
}
