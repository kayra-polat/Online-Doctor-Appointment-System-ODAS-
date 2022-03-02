package controllers

import (
	"cmpe313-odas-server/models"
	"cmpe313-odas-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 16); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		user.PasswordHash = string(hashedPassword)
	}

	db := models.GetDatabaseInstance()
	res := db.Create(&user)

	if err := res.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user.PasswordHash = ""
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(ctx *gin.Context) {
	var auth models.Login
	var user models.User

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email and/or password fields are missing"})
		ctx.Abort()
	}

	db := models.GetDatabaseInstance()
	res := db.Where("email = ?", auth.Email).First(&user)
	if err := res.Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unmatched email and/or password"})
		ctx.Abort()
		return
	}

	pwdMatchErr := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(auth.Password))
	if pwdMatchErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unmatched email and/or password"})
		ctx.Abort()
		return
	}

	token, err := utils.GenerateToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"data": token})
}
