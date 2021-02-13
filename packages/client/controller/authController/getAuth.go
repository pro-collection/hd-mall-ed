package authController

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/authModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
	"strings"
)

func GetAuth(c *gin.Context) {
	var err error
	api := app.ApiFunction{C: c}
	auth := authStruct{}
	_ = c.ShouldBindJSON(&auth)

	// 参数校验
	_, err = govalidator.ValidateStruct(auth)
	if err != nil {
		api.ResFailMessage(e.InvalidParams, strings.Split(err.Error(), ";")[0])
		return
	}

	data := make(map[string]interface{})
	code := e.FailAuthToken
	id, err := authModel.CheckAuth(auth.Name, auth.Password)
	if err != nil {
		api.ResFailMessage(e.FailAuthCheckToken, err.Error())
		return
	}

	if id > 0 {
		token, claims, err := jwtUtil.GenerateToken(auth.Name, auth.Password, id)

		if err != nil {
			code = e.FailAuthTokenCreate
		} else {
			data["token"] = token

			// 设置 cookie
			api.SetCookie("token", token, int(claims.ExpiresAt))
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
