package jwtUtil

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"hd-mall-ed/packages/client/config"
	"hd-mall-ed/packages/client/models/authModel"
	"hd-mall-ed/packages/client/models/userModel"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, id int) (string, Claims, error) {
	// 这个地方可以考虑通过密码动态授权
	var jwtSecret = []byte(config.AppConfig.JwtSecret + password)

	nowTime := time.Now()
	expireTime := nowTime.Add(config.DatabaseConfig.RedisExpire)

	claims := Claims{
		username,
		password,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hd-client",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, claims, err
}

func ParseToken(token string) (*Claims, error) {
	id := handleGetIdHelper(token)

	user := authModel.GetAuthById(id)

	// 通过密码动态授权
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JwtSecret + user.Password), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 通过 token 获取用户信息
func ParseTokenGetUser(token string) (userModel.User, error) {
	id := handleGetIdHelper(token)

	user := authModel.GetAuthById(id)
	if user.ID > 0 {
		return user, nil
	}
	return user, errors.New("获取用户信息失败")
}
