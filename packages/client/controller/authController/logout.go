package authController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/config/cache"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
	"hd-mall-ed/packages/client/pkg/utils/jwtUtil"
	"strconv"
)

// 退出登录状态
func Logout(c *gin.Context) {
	api := app.ApiFunction{C: c}
	// 获取 token
	token, err := c.Cookie("token")
	if err != nil {
		api.ResFail(e.FailLoginStatus)
		return
	}

	// 获取 user.id
	user, err2 := jwtUtil.ParseTokenGetUser(token)
	if err2 != nil {
		api.ResFailMessage(e.Fail, err2.Error())
		return
	}
	userId := strconv.Itoa(int(user.ID))

	// 移除 cookie
	api.RemoveCookie("token")

	// 移除 cache
	err = cache.Manager.Delete(userId)
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
