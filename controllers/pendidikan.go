package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetPendidikan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetPendidikan()
	if err != nil {
		return handleError(c, err, "Failed to retrieve pendidikan")
	}

	return sendResponse(c, fiber.StatusOK, true, "All pendidikan retrieved successfully", result)
}
func AddPendidikan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Parse body request for new Pendidikan
	newKategori := new(models.Pendidikan)
	err = c.BodyParser(newKategori)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.CreatePendidikan(newKategori.Name, newKategori.Description)
	if err != nil {
		return handleError(c, err, "Failed to add pendidikan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Pendidikan added successfully", result)
}

func UpdatePendidikan(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse body request for the updated Pendidikan
	newTask := new(models.Pendidikan)
	err = c.BodyParser(newTask)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdatePendidikan(newTask.Name, newTask.Description, id)
	if err != nil {
		return handleError(c, err, "Failed to update pendidikan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Pendidikan updated successfully", result)
}

func DeletePendidikan(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err = database.DeletePendidikan(id)
	if err != nil {
		return handleError(c, err, "Failed to delete pendidikan")
	}

	return sendResponse(c, fiber.StatusOK, true, "Pendidikan deleted successfully", nil)
}
