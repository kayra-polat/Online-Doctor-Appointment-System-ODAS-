package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID         uint64    `json:"userID" gorm:"primary_key;auto_increment;not_null"`
	Name           string    `json:"name" binding:"required"`
	Surname        string    `json:"surname" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Username       string    `json:"username" binding:"required"`
	Gender         string    `json:"gender" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Address        string    `json:"address" binding:"required"`
	DateOfBirth    time.Time `json:"dateOfBirth" binding:"required"`
	PasswordHash   string    `json:"passwordHash,omitempty" binding:"required"`
	Title          string    `json:"title"`
	Expertise      string    `json:"expertise"`
	Hospital       string    `json:"hospital"`
	Room           string    `json:"room"`
	PrivilegeLevel uint      `json:"privilegeLevel"`
	Role           string    `json:"role" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if r := u.Role; r == "Doctor" || r == "Patient" || r == "Admin" {
		switch r {
		case "Doctor":
			{
				if u.Title == "" || u.Expertise == "" || u.Hospital == "" || u.Room == "" {
					err = errors.New("required fields are empty: doctor")
					return
				}
			}
		case "Admin":
			{
				if u.PrivilegeLevel == 0 {
					err = errors.New("required fields are empty: admin")
					return
				}
			}

		}

		return
	}

	err = errors.New("user role is not valid")
	return
}
