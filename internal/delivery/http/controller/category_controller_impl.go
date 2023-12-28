package controller

import (
	"go-crud-v2/internal/model"
	"go-crud-v2/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryControllerImpl struct {
	UseCase usecase.CategoryUseCase
	Log     *logrus.Logger
}

func NewCategoryController(usecase usecase.CategoryUseCase, log *logrus.Logger) CategoryController {
	return &CategoryControllerImpl{
		UseCase: usecase,
		Log:     log,
	}
}

func (c *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCategoryRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parser body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateCategoryRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parser body")
		return fiber.ErrBadRequest
	}

	paramId, _ := strconv.Atoi(ctx.Params("categoryId"))
	request.ID = paramId

	response, err := c.UseCase.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	request := new(model.DeleteCategoryRequest)

	paramId, _ := strconv.Atoi(ctx.Params("categoryId"))
	request.ID = paramId
	if err := c.UseCase.Delete(ctx.Context(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete category")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *CategoryControllerImpl) Get(ctx *fiber.Ctx) error {
	request := new(model.GetCategoryRequest)

	paramId, _ := strconv.Atoi(ctx.Params("categoryId"))
	request.ID = paramId

	category, err := c.UseCase.Get(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: category})
}

func (c *CategoryControllerImpl) List(ctx *fiber.Ctx) error {
	categories, err := c.UseCase.List(ctx.Context())
	if err != nil {
		c.Log.WithError(err).Error("failed to get all category")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CategoryResponse]{Data: categories})
}
