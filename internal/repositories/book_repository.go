package repositories

import (
	"bookshelf/internal/entities"

	"gorm.io/gorm"
)

func GetAll(db *gorm.DB, entity *[]entities.Book) error {
	err := db.Find(&entity).Error

	return err
}

func GetById(db *gorm.DB, entity *entities.Book, id string) error {
	err := db.Find(&entity, id).Error

	return err
}

func Create(db *gorm.DB, entity *entities.Book) error {
	err := db.Create(&entity).Error
	return err
}

func Update(db *gorm.DB, entity *entities.Book, id string) error {
	err := db.Where("id = ?", id).Update("books", entity).Error
	return err
}

func Delete(db *gorm.DB, entity *entities.Book, id string) error {
	err := db.Where("id = ?", id).Delete(&entity).Error
	return err
}
