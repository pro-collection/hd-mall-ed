package userController

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/authModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"hd-mall-ed/packages/client/pkg/utils"
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
	id, _ := authModel.CheckAuth(auth.Username, auth.Password)
	if id > 0 {
		token, claims, err := utils.GenerateToken(auth.Username, auth.Password, id)

		if err != nil {
			code = e.FailAuthTokenCreate
		} else {
			data["token"] = token

			// 设置 cookie
			c.SetCookie("token", token, int(claims.ExpiresAt), "", "", false, false)
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
