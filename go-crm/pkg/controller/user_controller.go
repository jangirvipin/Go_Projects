package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jangirvipin/go-crm/pkg/model"
	"strconv"
)

func GetLeads(c *fiber.Ctx) error {
	leads, err := model.GetLeads()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch leads",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Leads fetched successfully",
		"data":    leads,
	})
}

func GetLeadsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID is required",
			"data":    nil,
		})
	}

	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
			"data":    nil,
		})
	}

	lead, err := model.GetLead(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch lead",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Lead fetched successfully",
		"data":    lead,
	})
}

func CreateLead(c *fiber.Ctx) error {
	lead := model.Lead{}
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	result := lead.Create()

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Lead created successfully",
		"data":    result,
	})
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID is required",
			"data":    nil,
		})
	}

	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID format",
			"data":    nil,
		})
	}

	result, err := model.DeleteLead(ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete lead",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Lead deleted successfully",
		"data":    result,
	})
}

func UpdateLead(c *fiber.Ctx) error {
	lead := model.Lead{}
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	id := c.Params("id")

	ID, _ := strconv.ParseInt(id, 0, 0)

	existingLead, err := model.GetLead(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch lead",
			"data":    nil,
		})
	}

	if lead.Name != "" {
		existingLead.Name = lead.Name
	}
	if lead.Email != "" {
		existingLead.Email = lead.Email
	}
	if lead.Phone != 0 {
		existingLead.Phone = lead.Phone
	}
	if lead.Company != "" {
		existingLead.Company = lead.Company
	}

	err = model.UpdateLead(existingLead)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update lead",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Lead updated successfully",
		"data":    existingLead,
	})
}
