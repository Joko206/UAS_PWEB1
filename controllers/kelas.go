package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	result, err := database.GetKelas()
	if err != nil {
		return handleError(c, err, "Failed to retrieve classes")
	}

	return sendResponse(c, fiber.StatusOK, true, "All classes retrieved successfully", result)
}

func AddKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	newKelas := new(models.Kelas)
	if err := c.BodyParser(newKelas); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.CreateKelas(newKelas.Name, newKelas.Description, user.ID)
	if err != nil {
		return handleError(c, err, "Failed to add class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class added successfully", result)
}

func UpdateKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	newTask := new(models.Kelas)
	if err := c.BodyParser(newTask); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdateKelas(newTask.Name, newTask.Description, id)
	if err != nil {
		return handleError(c, err, "Failed to update class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class updated successfully", result)
}

func DeleteKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err := database.DeleteKelas(id)
	if err != nil {
		return handleError(c, err, "Failed to delete class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class deleted successfully", nil)
}

// JoinKelasByCode allows user to join class using join code
func JoinKelasByCode(c *fiber.Ctx) error {
	// Authenticate the user
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	var requestData struct {
		JoinCode string `json:"join_code"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	if requestData.JoinCode == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Join code is required", nil)
	}

	// Join class using join code
	err = database.JoinKelasByCode(user.ID, requestData.JoinCode)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, err.Error(), nil)
	}

	// Get the class info to return
	kelas, err := database.GetKelasByJoinCode(requestData.JoinCode)
	if err != nil {
		return sendResponse(c, fiber.StatusInternalServerError, false, "Failed to get class info", nil)
	}

	return sendResponse(c, fiber.StatusOK, true, "Successfully joined class", kelas)
}
