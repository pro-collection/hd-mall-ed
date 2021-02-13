package jwtUtil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"hd-mall-ed/packages/admin/models/adminAuthModel"
	"hd-mall-ed/packages/admin/models/adminUserModel"
	"hd-mall-ed/packages/common/config"
	"time"
)

func AdminGenerateToken(username, password string, id int) (string, Claims, error) {
	// 这个地方可以考虑通过密码动态授权
	var jwtSecret = []byte(config.AppConfig.AdminJwtSecret + password)

	nowTime := time.Now()
	expireTime := nowTime.Add(config.DatabaseConfig.RedisExpire)

	claims := Claims{
		username,
		password,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hd-admin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, claims, err
}

func AdminParseToken(token string) (*Claims, error) {
	id := handleGetIdHelper(token)

	// 获取token
	user := adminAuthModel.GetAuthById(id)

	// 通过密码动态授权
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.AdminJwtSecret + user.Password), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 通过 token 获取用户信息
func AdminParseTokenGetUser(token string) (adminUserModel.AdminUser, error) {
	id := handleGetIdHelper(token)

	// 获取 admin 用户信息
	user := adminAuthModel.GetAuthById(id)
	if user.ID > 0 {
		return user, nil
	}
	return user, errors.New("获取用户信息失败")
}
