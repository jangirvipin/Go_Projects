package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-urlShortner/controller"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/:key", controller.RedirectHandler)
	app.Post("/shorten", controller.RegisterHandler)
}
