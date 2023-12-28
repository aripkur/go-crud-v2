package repository

import (
	"go-crud-v2/internal/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(tx *gorm.DB, category *entity.Category) error
	Update(tx *gorm.DB, category *entity.Category) error
	Delete(tx *gorm.DB, category *entity.Category) error
	FindById(tx *gorm.DB, category *entity.Category, categoryId int) error
	FindAll(tx *gorm.DB) ([]entity.Category, error)
}
