package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewAuthMiddleware(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Get("X-API-Key") != "RAHASIA" {
			log.Error("unathorized")
			return fiber.ErrUnauthorized
		}
		return ctx.Next()
	}

}
