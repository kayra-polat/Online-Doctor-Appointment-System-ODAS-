package utils

import (
	"cmpe313-odas-server/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const Secret = "verySecretKey"

func GenerateToken(u *models.User) (models.Token, error) {
	var newToken models.Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userID"] = u.UserID
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return newToken, errors.New(fmt.Sprintf("token generation failed: %s", err.Error()))
	}

	newToken.TokenString = tokenString
	u.PasswordHash = ""
	newToken.Profile = *u
	return newToken, nil
}
