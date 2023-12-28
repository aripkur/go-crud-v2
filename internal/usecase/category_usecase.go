package usecase

import (
	"context"
	"go-crud-v2/internal/model"
)

type CategoryUseCase interface {
	Create(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error)
	Update(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error)
	Delete(ctx context.Context, request *model.DeleteCategoryRequest) error
	Get(ctx context.Context, request *model.GetCategoryRequest) (*model.CategoryResponse, error)
	List(ctx context.Context) ([]model.CategoryResponse, error)
}
