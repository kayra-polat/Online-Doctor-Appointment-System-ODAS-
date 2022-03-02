package controllers

import (
	"cmpe313-odas-server/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadProfile(ctx *gin.Context) {
	var u models.User
	//var token models.Token

	//_ = ctx.ShouldBindJSON(&token)

	uId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid field 'id'"})
		ctx.Abort()
		return
	}

	//if token.UserID != uId {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	//	return
	//}

	u.UserID = uId
	userID, _ := ctx.Get("userID")

	if u.UserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{})
		ctx.Abort()
		return
	}

	db := models.GetDatabaseInstance()
	res := db.First(&u)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			ctx.Abort()
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
		}
		return
	}

	u.PasswordHash = ""
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"data": u})
}
