package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/client/database/tableModel"
	"hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"log"
	"strconv"
)

// 查询用户信息， 需要登录 和 权限验证
// 只能查询自己的用户信息
func GetUserInfo(c *gin.Context) {
	model := &userModel.User{}
	api := app.ApiFunction{C: c}
	baseUser := &tableModel.BaseUser{}

	userId := c.DefaultQuery("id", "")
	if userId == "" {
		api.ResFail(e.InvalidParams)
		return
	}

	id, _ := strconv.Atoi(userId)

	// 问题类型转换
	user, err := model.FindUserById(uint(id))
	if err != nil {
		log.Println(err.Error())
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	_ = deepcopier.Copy(user).To(baseUser)
	api.Response(baseUser)
}
