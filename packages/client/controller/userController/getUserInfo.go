package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
)

// 查询用户信息， 需要登录 和 权限验证
// 只能查询自己的用户信息
func GetUserInfo(c *gin.Context) {
	api := app.ApiFunction{C: c}

	baseUser := &tableModel.BaseUser{}
	var err error
	userMapper := &userModel.User{}

	userid := uint(api.GetUserId())

	user, err := userMapper.FindUserById(userid)

	if err != nil {
		api.ResFail(e.InvalidParams)
		return
	}

	_ = deepcopier.Copy(user).To(baseUser)
	api.Response(baseUser)
}
