package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      viper.GetString("app.name"),
		ErrorHandler: NewErrorHandle(),
		Prefork:      viper.GetBool("web.prefork"),
	})

	return app
}

func NewErrorHandle() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var message any

		code := fiber.StatusInternalServerError
		message = err.Error()

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		if e, ok := err.(validator.ValidationErrors); ok {
			var validateMessages []string
			for _, errorField := range e {
				validateMessages = append(validateMessages, errorField.Field()+" is "+errorField.Tag())
			}
			code = 402
			message = validateMessages
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": message,
		})
	}
}
