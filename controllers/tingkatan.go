package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetTingkatan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetTingkatan()
	if err != nil {
		return handleError(c, err, "Failed to retrieve Tingkatan")
	}

	return sendResponse(c, fiber.StatusOK, true, "All Tingkatan retrieved successfully", result)
}

func AddTingkatan(c *fiber.Ctx) error {

	// Authenticate the user using the JWT token

	// Parse body request for new Tingkatan
	newTingkatan := new(models.Tingkatan)
	err := c.BodyParser(newTingkatan)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Create Tingkatan
	result, err := database.CreateTingkatan(newTingkatan.Name, newTingkatan.Description)
	if err != nil {
		return handleError(c, err, "Failed to add Tingkatan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Tingkatan added successfully", result)
}

func UpdateTingkatan(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse body request for updated Tingkatan
	newTingkatan := new(models.Tingkatan)
	err := c.BodyParser(newTingkatan)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Update Tingkatan
	result, err := database.UpdateTingkatan(newTingkatan.Name, newTingkatan.Description, id)
	if err != nil {
		return handleError(c, err, "Failed to update Tingkatan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Tingkatan updated successfully", result)
}

func DeleteTingkatan(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Delete Tingkatan
	err := database.DeleteTingkatan(id)
	if err != nil {
		return handleError(c, err, "Failed to delete Tingkatan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Tingkatan deleted successfully", nil)
}
