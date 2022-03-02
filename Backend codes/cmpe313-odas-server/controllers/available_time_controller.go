package controllers

import (
	"cmpe313-odas-server/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAvailableTime(ctx *gin.Context) {
	var at models.AvailableTime

	if err := ctx.ShouldBindJSON(&at); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	doctorID, _ := ctx.Get("userID")
	if at.DoctorID != doctorID {
		log.Println(at, doctorID)
		ctx.JSON(http.StatusForbidden, gin.H{})
		ctx.Abort()
		return
	}

	db := models.GetDatabaseInstance()
	res := db.Create(&at)

	if err := res.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusCreated, gin.H{"data": at})
}

func ReadAllAvailableTime(ctx *gin.Context) {
	var at []models.AvailableTime
	//fat := models.FilterableAvailableTime()

	//// TODO: Fix filtering system: filters does  not work for /get available time
	//
	//params := ctx.Params
	//for _, param := range params {
	//	for _, key := range fat.ValidKeys {
	//		if param.Key == key {
	//			fat.Filters[key] = param.Value
	//		}
	//	}
	//}

	db := models.GetDatabaseInstance()
	//doctorID, _ := ctx.Get("userID")
	res := db.Find(&at)

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

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"data": at, "length": res.RowsAffected})
}

func DeleteAvailableTime(ctx *gin.Context) {
	var at models.AvailableTime

	atId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid field 'id'"})
		ctx.Abort()
		return
	}

	at.AvailableTimesID = atId

	db := models.GetDatabaseInstance()
	doctorID, _ := ctx.Get("userID")
	res := db.Where("doctorID = ?", doctorID).Delete(&at)

	if err := res.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusNoContent, gin.H{})
}
