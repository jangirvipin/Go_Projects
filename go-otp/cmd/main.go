package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-otp/api/routes"
	"log"
)

func main() {
	app := fiber.New()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
