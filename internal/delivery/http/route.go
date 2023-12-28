package http

import (
	"go-crud-v2/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	CategoryController controller.CategoryController
	AuthMiddleware     fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)

	c.App.Get("/api/categories/:categoryId", c.CategoryController.Get)
	c.App.Get("/api/categories", c.CategoryController.List)
	c.App.Post("/api/categories", c.CategoryController.Create)
	c.App.Delete("/api/categories/:categoryId", c.CategoryController.Delete)
	c.App.Put("/api/categories/:categoryId", c.CategoryController.Update)
}
