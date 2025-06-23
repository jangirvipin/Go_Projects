package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-urlShortner/redis"
)

type Body struct {
	URL string `json:"url"`
}

func RedirectHandler(c *fiber.Ctx) error {
	key := c.Params("key")
	if key == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Key parameter is required",
		})
	}

	rClient := redis.NewRedisClient()

	url, err := rClient.GetHandler(key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Redirect(url, fiber.StatusFound)
}

func RegisterHandler(c *fiber.Ctx) error {
	var body Body
	if err := c.BodyParser(&body); err != nil {
		c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Bad request body",
		})
	}
	url := body.URL

	if url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "URL is required",
		})
	}

	rClient := redis.NewRedisClient()
	shortKey, err := rClient.SetHandler(url)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	var finalURL = url + "/" + shortKey
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"url":     finalURL,
	})
}
