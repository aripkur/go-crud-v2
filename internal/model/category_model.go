package model

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetCategoryRequest struct {
	ID int `json:"id" validate:"required"`
}
