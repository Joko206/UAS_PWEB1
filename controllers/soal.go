package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	result, err := database.GetSoal()
	if err != nil {
		return handleError(c, err, "Failed to retrieve soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "All soal retrieved successfully", result)
}

func AddSoal(c *fiber.Ctx) error {

	// Authenticate the user using the JWT token

	// Parse body request for new Soal
	newSoal := new(models.Soal)
	err := c.BodyParser(newSoal)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Create Soal
	result, err := database.CreateSoal(newSoal.Question, newSoal.Options, newSoal.Correct_answer, newSoal.Kuis_id)
	if err != nil {
		return handleError(c, err, "Failed to add soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal added successfully", result)
}

func UpdateSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse body request for updated Soal
	newSoal := new(models.Soal)
	err := c.BodyParser(newSoal)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Update Soal
	result, err := database.UpdateSoal(newSoal.Question, newSoal.Options, newSoal.Correct_answer, newSoal.Kuis_id, id)
	if err != nil {
		return handleError(c, err, "Failed to update soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal updated successfully", result)
}
func DeleteSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Delete Soal
	err := database.DeleteSoal(id)
	if err != nil {
		return handleError(c, err, "Failed to delete soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal deleted successfully", nil)
}
func GetSoalByKuisID(c *fiber.Ctx) error {
	// Ambil kuis_id dari parameter request
	kuisID := c.Params("kuis_id")

	// Cek apakah kuis_id valid
	var kuis models.Kuis
	err := database.DB.First(&kuis, kuisID).Error
	if err != nil {
		return sendResponse(c, fiber.StatusNotFound, false, "Kuis not found", nil)
	}
	var soal []models.Soal
	err = database.DB.Where("kuis_id = ?", kuisID).Find(&soal).Error
	if err != nil {
		return sendResponse(c, fiber.StatusInternalServerError, false, "Failed to fetch questions", nil)
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal retrieved successfully", soal)
}
