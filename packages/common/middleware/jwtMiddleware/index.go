package jwtMiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"hd-mall-ed/packages/common/pkg/utils/jwtUtil"
	"time"
)

// 鉴权中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := app.ApiFunction{C: c}

		var code int

		code = e.Success

		// 需要从 cookie 中获取
		token, err := c.Cookie("token")
		if err != nil {
			code = e.FailAuthCheckToken
		}

		if funk.IsEmpty(token) {
			code = e.FailAuthCheckToken
		} else {
			claims, err := jwtUtil.ParseToken(token)
			if err != nil {
				code = e.FailAuthCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 这个情况是超时了
				code = e.FailAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			api.ResFail(code)
			c.Abort()
			return
		}

		c.Next()
	}
}

// admin 鉴权中间件
func AdminJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := adminApp.ApiInit(c)

		var code int

		code = e.Success

		// 需要从 cookie 中获取
		token, err := c.Cookie("admin-token")
		if err != nil {
			code = e.FailAuthCheckToken
		}

		if funk.IsEmpty(token) {
			code = e.FailAuthCheckToken
		} else {
			claims, err := jwtUtil.AdminParseToken(token)
			if err != nil {
				code = e.FailAuthCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 这个情况是超时了
				code = e.FailAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			api.ResFail(code)
			c.Abort()
			return
		}

		c.Next()
	}
}
