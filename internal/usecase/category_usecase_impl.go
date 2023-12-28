package usecase

import (
	"context"
	"go-crud-v2/internal/entity"
	"go-crud-v2/internal/model"
	"go-crud-v2/internal/model/converter"
	"go-crud-v2/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryUseCaseImpl struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryUseCaseImpl) Create(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	category := new(entity.Category)
	category.Name = request.Name

	if err := c.CategoryRepository.Save(tx, category); err != nil {
		c.Log.WithError(err).Error("failed to save category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction category")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCaseImpl) Update(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find category")
		return nil, fiber.ErrNotFound
	}

	category.Name = request.Name

	if err := c.CategoryRepository.Update(tx, category); err != nil {
		c.Log.WithError(err).Error("failed to delete category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction category")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCaseImpl) Delete(ctx context.Context, request *model.DeleteCategoryRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	category := new(entity.Category)
	category.ID = request.ID

	if err := c.CategoryRepository.Delete(tx, category); err != nil {
		c.Log.WithError(err).Error("failed to delete category")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CategoryUseCaseImpl) Get(ctx context.Context, request *model.GetCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to get category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCaseImpl) List(ctx context.Context) ([]model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	categories, err := c.CategoryRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("failed to get all category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *converter.CategoryToResponse(&category)
	}

	return responses, nil

}
