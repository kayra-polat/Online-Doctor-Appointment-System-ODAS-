package middlewares

import (
	"cmpe313-odas-server/models"
	"cmpe313-odas-server/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

func ValidateAuthorization(permittedRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization header"})
			ctx.Abort()
			return
		} else {
			tokenHeader = tokenHeader[7:]
		}

		var claims models.UserClaims
		token, err := jwt.ParseWithClaims(tokenHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("jwt parsing error")
			}
			return []byte(utils.Secret), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid and/or expired"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
			if role := claims.Role; role == permittedRole || permittedRole == "Both" {
				log.Println(claims, "\n")
				ctx.Set("userID", claims.UserID)
				ctx.Set("role", claims.Role)
				ctx.Next()
				return
			}

			ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden role"})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusForbidden, gin.H{"error": "claims are not valid or something else"})
		ctx.Abort()
		return
	}
}
