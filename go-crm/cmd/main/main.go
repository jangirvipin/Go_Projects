package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jangirvipin/go-crm/pkg/routes"
)

func main() {
	app := fiber.New()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
