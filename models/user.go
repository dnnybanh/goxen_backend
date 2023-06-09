package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	FirstName string `json:"firstName" gorm:"size:100"`
	LastName  string `json:"lastName" gorm:"size:100"`
	UserType  string `json:"userType" gorm:"size:100"`
}

// BeforeCreate is a GORM hook that will be triggered every time a new user is stored in the database
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that will be triggered every time an existing user is updated in the database
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
