package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/admin/models/adminUserModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"log"
)

// @title           注册 admin 账户
// @description
// @auth            晴小篆  331393627@qq.com
// @param   		adminUserModel.AdminUser
// @return
func Register(c *gin.Context) {
	user := &adminUserModel.AdminUser{}
	api := adminApp.ApiInit(c)
	var err error

	// 绑定参数
	err = c.ShouldBindJSON(user)
	if err != nil {
		api.ResFail(e.FailBindJson)
		return
	}

	// 确认用户名参数存在
	if funk.IsEmpty(user.Name) {
		api.ResFail(e.InvalidParams)
		return
	}

	// 校验参数
	if api.ValidateHasError(user) {
		return
	}

	// 判定是否存在当前用户信息
	if handleCheckUserNameIsExistHelper(user, api) {
		return
	}

	// 更新用户信息
	if funk.IsEmpty(user.Role) {
		user.Role = "2" // 普通管理员
	}

	// 创建用户
	err = user.CreateUser()
	if err != nil {
		log.Println(err.Error())
		api.ResFail(e.CreateUserFail)
		return
	}

	api.ResponseNoData()
}
