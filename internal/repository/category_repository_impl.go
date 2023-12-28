package repository

import (
	"go-crud-v2/internal/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(tx *gorm.DB, category *entity.Category) error {
	return tx.Create(category).Error
}

func (c *CategoryRepositoryImpl) Update(tx *gorm.DB, category *entity.Category) error {
	return tx.Save(category).Error
}

func (c *CategoryRepositoryImpl) Delete(tx *gorm.DB, category *entity.Category) error {
	return tx.Delete(category).Error
}

func (c *CategoryRepositoryImpl) FindById(tx *gorm.DB, category *entity.Category, categoryId int) error {
	return tx.Where("id = ?", categoryId).First(category).Error

}

func (c *CategoryRepositoryImpl) FindAll(tx *gorm.DB) ([]entity.Category, error) {
	var categories []entity.Category
	if err := tx.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
