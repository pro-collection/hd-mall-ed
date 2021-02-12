package jwtUtil

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}
