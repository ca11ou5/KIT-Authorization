package postgres

import (
	"Authorization_Service/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	db *gorm.DB
}

func InitDB(host, user, password, dbname, port, sslmode string) Handler {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	return Handler{db}
}
