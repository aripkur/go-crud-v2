package config

import (
	"go-crud-v2/internal/delivery/http"
	"go-crud-v2/internal/delivery/http/controller"
	"go-crud-v2/internal/delivery/http/middleware"

	"go-crud-v2/internal/repository"
	"go-crud-v2/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	Log      *logrus.Logger
	App      *fiber.App
	Validate *validator.Validate
	Viper    *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	categoryRepository := repository.NewCategoryRepository()
	categoryUseCase := usecase.NewCategoryUseCase(config.DB, config.Log, config.Validate, categoryRepository)
	categoryController := controller.NewCategoryController(categoryUseCase, config.Log)

	authMiddleware := middleware.NewAuthMiddleware(config.Log)

	routeConfig := http.RouteConfig{
		App:                config.App,
		CategoryController: categoryController,
		AuthMiddleware:     authMiddleware,
	}

	routeConfig.Setup()
}
