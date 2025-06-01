package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user has permission (admin or teacher)
	if user.Role != "admin" && user.Role != "teacher" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Only admin and teacher can view all questions", nil)
	}

	result, err := database.GetSoal()
	if err != nil {
		return handleError(c, err, "Failed to retrieve soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "All soal retrieved successfully", result)
}

func AddSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user has permission (admin or teacher)
	if user.Role != "admin" && user.Role != "teacher" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Only admin and teacher can add questions", nil)
	}

	// Parse body request for new Soal
	newSoal := new(models.Soal)
	err = c.BodyParser(newSoal)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Validate required fields
	if newSoal.Question == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Question cannot be empty", nil)
	}
	if newSoal.CorrectAnswer == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Correct answer cannot be empty", nil)
	}
	if newSoal.KuisID == 0 {
		return sendResponse(c, fiber.StatusBadRequest, false, "Quiz ID cannot be empty", nil)
	}

	// Create Soal
	result, err := database.CreateSoal(newSoal.Question, newSoal.Options, newSoal.CorrectAnswer, newSoal.KuisID)
	if err != nil {
		return handleError(c, err, "Failed to add soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal added successfully", result)
}

func UpdateSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user has permission (admin or teacher)
	if user.Role != "admin" && user.Role != "teacher" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Only admin and teacher can update questions", nil)
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse body request for updated Soal
	newSoal := new(models.Soal)
	err = c.BodyParser(newSoal)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Update Soal
	result, err := database.UpdateSoal(newSoal.Question, newSoal.Options, newSoal.CorrectAnswer, newSoal.KuisID, id)
	if err != nil {
		return handleError(c, err, "Failed to update soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal updated successfully", result)
}
func DeleteSoal(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user has permission (admin or teacher)
	if user.Role != "admin" && user.Role != "teacher" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Only admin and teacher can delete questions", nil)
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Delete Soal
	err = database.DeleteSoal(id)
	if err != nil {
		return handleError(c, err, "Failed to delete soal")
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal deleted successfully", nil)
}
func GetSoalByKuisID(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	kuisID := c.Params("kuis_id")
	if kuisID == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Kuis ID cannot be empty", nil)
	}

	// Use existing database connection
	db := database.MustGetDB()

	// Cek apakah kuis_id valid
	var kuis models.Kuis
	err = db.First(&kuis, kuisID).Error
	if err != nil {
		return sendResponse(c, fiber.StatusNotFound, false, "Kuis not found", nil)
	}

	// Get questions for the quiz - using the correct column name
	var soal []models.Soal
	err = db.Where("kuis_id = ?", kuisID).Find(&soal).Error
	if err != nil {
		return sendResponse(c, fiber.StatusInternalServerError, false, "Failed to fetch questions", nil)
	}

	return sendResponse(c, fiber.StatusOK, true, "Soal retrieved successfully", soal)
}
