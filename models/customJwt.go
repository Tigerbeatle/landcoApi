package models

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	UserUUID string `json:"userUUID"`
	jwt.StandardClaims
}
