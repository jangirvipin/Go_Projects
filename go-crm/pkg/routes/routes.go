package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-crm/pkg/controller"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/v1/leads", controller.GetLeads)
	api.Post("/v1/leads", controller.CreateLead)
	api.Get("/v1/leads/:id", controller.GetLeadsByID)
	api.Delete("/v1/leads/:id", controller.DeleteLead)
	api.Put("/v1/leads/:id", controller.UpdateLead)
}
