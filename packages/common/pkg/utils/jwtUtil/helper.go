package jwtUtil

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
	"strings"
)

func handleGetIdHelper(token string) int64 {
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
	return id
}
