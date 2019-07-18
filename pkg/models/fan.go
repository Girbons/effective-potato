package models

import "github.com/jinzhu/gorm"

// Fan reprensets a Fan device
type Fan struct {
	gorm.Model
	Name   string `json:"name"`
	Pin    int    `gorm:"unique" json:"pin"`
	Status bool   `json:"status"`

	Temperature float32 `json:"temperature"`
}
