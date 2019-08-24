package models

import (
	"github.com/jinzhu/gorm"
)

// User represents the User
// with Username, Password, isAuthenticated
type User struct {
	gorm.Model

	Username        string `gorm:"unique" json:"Username"`
	Password        string `json:"-"` // omit the password from serialization
	IsAuthenticated bool   `json:"is_authenticated"`
}
