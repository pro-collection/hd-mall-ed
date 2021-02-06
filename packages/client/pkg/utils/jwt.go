package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
	"hd-mall-ed/packages/client/config"
	"hd-mall-ed/packages/client/models/authModel"
	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, id int) (string, error) {
	// 这个地方可以考虑通过密码动态授权
	var jwtSecret = []byte(config.AppConfig.JwtSecret + password)

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

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
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	payload := strings.Split(token, ".")

	// 获取token 的中间段信息
	bytes, e := jwt.DecodeSegment(payload[1])

	if e != nil {
		println(e.Error())
	}
	content := ""
	for i := 0; i < len(bytes); i++ {
		content += string(bytes[i])
	}

	id := gjson.Get(content, "id").Int()

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
