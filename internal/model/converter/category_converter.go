package converter

import (
	"go-crud-v2/internal/entity"
	"go-crud-v2/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
