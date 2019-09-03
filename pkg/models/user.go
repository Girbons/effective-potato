package models

import (
	"github.com/jinzhu/gorm"
)

// User represents a User
// with Username, Password
type User struct {
	gorm.Model

	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"` // omit
}
