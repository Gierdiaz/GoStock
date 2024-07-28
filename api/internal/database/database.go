package database

import (
	"log"

	"github.com/Gierdiaz/GoStock/api/config"
	"github.com/Gierdiaz/GoStock/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	config := config.LoadConfig()
	dsn := config.GetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.AutoMigrate(
		&models.User{},
	)

	DB = db

}