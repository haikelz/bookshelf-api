package configs

import (
	"bookshelf/internal/entities"
	"bookshelf/internal/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm() *gorm.DB {
	db, err := gorm.Open(postgres.Open(utils.Env().DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate
	err = db.AutoMigrate(&entities.Book{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	return db
}
