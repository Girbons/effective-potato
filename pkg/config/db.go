package config

import (
	"github.com/Girbons/effective-potato/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Light{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Fan{})
	db.AutoMigrate(&models.Room{})
}

// InitDB will initialize the database and returns a DB connection
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "potato.db")

	if err != nil {
		panic("failed to connect database")
	}

	migrateModels(db)
	return db
}
