package authController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/common/config/cache"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
)

func Logout(c *gin.Context) {
	api := adminApp.ApiInit(c)

	// 获取 token
	token, err := c.Cookie(adminTokenKey)
	if err != nil && funk.IsEmpty(token) {
		api.ResFail(e.FailLoginStatus)
		return
	}

	// 移除 cookie
	api.RemoveCookie(adminTokenKey)

	// 获取当前登录的用户信息
	user, err2 := jwtUtil.AdminParseTokenGetUser(token)
	if err2 != nil {
		api.ResFailMessage(e.Fail, err2.Error())
		return
	}

	// 移除 cache
	err = cache.Manager.Delete(cache.AminCacheKey(int(user.ID)))
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	api.ResponseNoData()
}
