package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/client/database/tableModel"
	. "hd-mall-ed/packages/client/models/userModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"log"
	"net/http"
	"strconv"
)

// 创建用户信息
func CreateUser(c *gin.Context) {
	var err error
	api := app.ApiFunction{C: c}
	user := &User{}

	// 绑定user信息， 以及参数校验
	if handleBindUserHasErrorHelper(user, c) {
		return
	}

	if user.Name == "" {
		api.ResFail(e.InvalidParams)
		return
	}

	// 判断用户是否存在
	if handleCheckUserExistHelper(user, c) {
		return
	}

	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    e.CreateUserFail,
			"message": e.GetMsg(e.CreateUserFail),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    e.Success,
		"message": e.GetMsg(e.Success),
	})
}

func UpdateUser(c *gin.Context) {
	var err error
	api := app.ApiFunction{C: c}
	user := &User{}

	// 绑定参数
	if handleBindUserHasErrorHelper(user, c) {
		return
	}

	// 校验 id 是否存在
	if handleCheckUserIdNotExistHelper(user, c) {
		return
	}

	// 校验用户名是否重复
	if handleCheckUserExistHelper(user, c) {
		return
	}

	// save 是全量更新
	updateMap := make(map[string]interface{})
	updateMap["password"] = user.Password

	err = user.Update(updateMap)
	if err != nil {
		log.Println(err.Error())
		api.ResFail(e.Fail)
		return
	}

	api.ResponseNoData()
}

// 查询用户信息， 需要登录 和 权限验证
// 只能查询自己的用户信息
func GetUserInfo(c *gin.Context) {
	userModal := &User{}
	api := app.ApiFunction{C: c}
	baseUser := &tableModel.BaseUser{}

	userId := c.DefaultQuery("id", "")
	if userId == "" {
		api.ResFail(e.InvalidParams)
		return
	}

	id, _ := strconv.Atoi(userId)

	// 问题类型转换
	user, err := userModal.FindUserById(uint(id))
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	_ = deepcopier.Copy(user).To(baseUser)
	api.Response(baseUser)
}
