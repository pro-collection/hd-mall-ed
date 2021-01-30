package userController

import (
	"github.com/gin-gonic/gin"
	. "hd-mall-ed/packages/client/models/userModel"
	"net/http"
)

// 创建用户信息
func CreateUser(c *gin.Context) {
	var err error
	user := &User{}

	// 绑定user信息， 以及参数校验
	if handleBindUserHasErrorHelper(user, c) {
		return
	}

	// 判断用户是否存在
	if handleCheckUserExistHelper(user, c) {
		return
	}

	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "创建用户失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建用户成功过",
	})
}
