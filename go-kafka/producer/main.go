package producer

import (
	"github.com/gofiber/fiber"
)

type Commnet struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	app.
}
