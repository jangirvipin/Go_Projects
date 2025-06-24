package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-otp/api/handlers"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/api/v1/otp", handlers.SendOtp)
	app.Post("/api/v1/verify", handlers.VerifyOtp)
}
