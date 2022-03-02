package controllers

import (
	"cmpe313-odas-server/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAppointment(ctx *gin.Context) {
	var ap models.Appointment

	if err := ctx.ShouldBindJSON(&ap); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	patientID, _ := ctx.Get("userID")
	if ap.PatientID != patientID {
		ctx.JSON(http.StatusForbidden, gin.H{})
		ctx.Abort()
		return
	}

	db := models.GetDatabaseInstance()
	res := db.Create(&ap)

	if err := res.Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusCreated, gin.H{"data": ap})
}

func ReadAllAppointments(ctx *gin.Context) {
	var ap []models.Appointment

	db := models.GetDatabaseInstance()
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("role")

	var res *gorm.DB
	if role == "Doctor" {
		res = db.Where("doctor_id = ?", userID).Find(&ap)
	} else {
		res = db.Where("patient_id = ?", userID).Find(&ap)
	}

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
	ctx.JSON(http.StatusOK, gin.H{"data": ap, "length": res.RowsAffected})
}

func ReadAppointment(ctx *gin.Context) {
	var ap models.Appointment

	apId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid field 'id'"})
		ctx.Abort()
		return
	}
	ap.AppointmentID = apId

	db := models.GetDatabaseInstance()
	patientID, _ := ctx.Get("userID")
	res := db.Where("patient_id = ?", patientID).First(&ap)

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
	ctx.JSON(http.StatusOK, gin.H{"data": ap})
}

func DeleteAppointment(ctx *gin.Context) {
	var ap models.Appointment

	apId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid field 'id'"})
		ctx.Abort()
		return
	}
	ap.AppointmentID = apId

	db := models.GetDatabaseInstance()
	patientID, _ := ctx.Get("userID")
	res := db.Where("patient_id = ?", patientID).Delete(&ap)

	if err := res.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusForbidden, gin.H{"Message": fmt.Sprintf("Appointment %d can't be deleted", ap.AppointmentID)})
		ctx.Abort()
		return
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{"Message": fmt.Sprintf("Appointment %d Cancelled Successfully", ap.AppointmentID)})
}
