package database

import (
	"fmt"
	"github.com/Brrocat/users-service/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Bogdan_20 dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическое создание таблиц
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}

	log.Println("Database tables created successfully")
	return db, nil
}
