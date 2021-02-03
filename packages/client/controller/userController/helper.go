package userController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/userModel"
	"net/http"
)

// 绑定user信息， 以及参数校验
func handleBindUserHasErrorHelper(user *userModel.User, c *gin.Context) bool {
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return true
	}
	return false
}

// 判断该用户是否 【不存在】， 更新用户的时候判定
// 不存在的时候， 返回true
// 存在用户的时候， 返回 false
func handleCheckUserIdNotExistHelper(user *userModel.User, c *gin.Context) bool {
	if user.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "缺少id",
		})
		return true
	}

	return false
}

// 判断用户是否存在
func handleCheckUserExistHelper(user *userModel.User, c *gin.Context) bool {
	resultUser, _ := user.FindUserByName(user.Name)
	if resultUser.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"name":    resultUser.Name,
			"message": "用户名已经存在",
		})
		return true
	}
	return false
}
