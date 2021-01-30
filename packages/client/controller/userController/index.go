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

	err = c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
	}
}
