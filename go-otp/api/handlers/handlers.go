package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-otp/api/otp"
)

func SendOtp(c *fiber.Ctx) error {

	var phone string
	if err := c.BodyParser(&phone); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	client := otp.NewOtpClient()

	otpCode, err := client.SendOtp(phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send OTP",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP sent successfully",
		"otp":     otpCode,
	})

}

func VerifyOtp(c *fiber.Ctx) error {
	payload := struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	client := otp.NewOtpClient()

	err, verified := client.VerifyOtp(payload.Phone, payload.Otp)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to verify OTP",
		})
	}
	if !verified {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid OTP",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP verified successfully",
	})
}
