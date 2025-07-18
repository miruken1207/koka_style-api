package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}
