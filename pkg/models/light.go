package models

import "github.com/jinzhu/gorm"

// Light represents a Light Device
type Light struct {
	gorm.Model

	Name   string `json:"name"`
	Pin    int    `gorm:"unique" json:"pin"`
	Status bool   `json:"status"`
}
