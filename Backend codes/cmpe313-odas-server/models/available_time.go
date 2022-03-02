package models

import (
	"time"
)

type AvailableTime struct {
	AvailableTimesID uint64    `json:"availableTimeID" gorm:"primary_key;auto_increment;not_null"`
	DoctorID         uint64    `json:"doctorID" binding:"required"`
	Date             time.Time `json:"date" binding:"required"`
	Hospital         string    `json:"hospital" binding:"required"`
	Room             string    `json:"room" binding:"required"`
}
