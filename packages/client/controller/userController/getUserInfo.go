package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"hd-mall-ed/packages/common/database/tableModel"
)

// 查询用户信息， 需要登录 和 权限验证
// 只能查询自己的用户信息
func GetUserInfo(c *gin.Context) {
	api := app.ApiFunction{C: c}
	baseUser := &tableModel.BaseUser{}
	var err error
	var user userModel.User

	user, err = api.GetUser()
	if err != nil {
		api.ResFail(e.InvalidParams)
		return
	}

	_ = deepcopier.Copy(user).To(baseUser)
	api.Response(baseUser)
}
