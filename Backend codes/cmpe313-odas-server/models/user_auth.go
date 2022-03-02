package models

import "github.com/golang-jwt/jwt"

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Profile     User   `json:"profile"`
	TokenString string `json:"tokenString"`
}

type UserClaims struct {
	UserID uint64 `json:"userID,omitempty"`
	Role   string `json:"role,omitempty"`
	jwt.StandardClaims
}
