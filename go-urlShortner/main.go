package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-urlShortner/routes"
	"log"
)

func main() {
	app := fiber.New()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
