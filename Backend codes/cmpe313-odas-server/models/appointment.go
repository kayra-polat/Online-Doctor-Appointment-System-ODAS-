package models

import (
	"time"
)

type Appointment struct {
	AppointmentID uint64    `json:"appointmentID" gorm:"primary_key;auto_increment;not_null"`
	Title         string    `json:"title" binding:"required"`
	PatientID     uint64    `json:"patientID" binding:"required"`
	DoctorID      uint64    `json:"doctorID" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Hospital      string    `json:"hospital" binding:"required"`
	Room          string    `json:"room" binding:"required"`
}
