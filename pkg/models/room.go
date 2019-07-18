package models

import "github.com/jinzhu/gorm"

// Represents a Room whith a Name and N Devices
type Room struct {
	gorm.Model

	Name   string   `json:"name"`
	Lights []*Light `json:"lights"`
	Fans   []*Fan   `json:"fans"`
}
