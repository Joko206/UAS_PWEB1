package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetKuis(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Get kuis that user can access (public + private from joined classes)
	result, err := database.GetKuisForUser(user.ID)
	if err != nil {
		return handleError(c, err, "Failed to retrieve quizzes")
	}

	return sendResponse(c, fiber.StatusOK, true, "Accessible quizzes retrieved successfully", result)
}

// GetAllKuis retrieves all kuis (admin only)
func GetAllKuis(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user is admin
	if user.Role != "admin" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Admin only.", nil)
	}

	result, err := database.GetKuis()
	if err != nil {
		return handleError(c, err, "Failed to retrieve quizzes")
	}

	return sendResponse(c, fiber.StatusOK, true, "All quizzes retrieved successfully", result)
}

func AddKuis(c *fiber.Ctx) error {
	// Authenticate user first
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Parse request body
	newKuis := new(models.Kuis)
	if err := c.BodyParser(newKuis); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Validate Kategori
	var kategori models.Kategori_Soal
	if err := db.First(&kategori, newKuis.Kategori_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Kategori ID", nil)
	}

	// Validate Tingkatan
	var tingkatan models.Tingkatan
	if err := db.First(&tingkatan, newKuis.Tingkatan_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Tingkatan ID", nil)
	}

	// Validate Kelas
	var kelas models.Kelas
	if err := db.First(&kelas, newKuis.Kelas_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Kelas ID", nil)
	}

	// Validate Pendidikan
	var pendidikan models.Pendidikan
	if err := db.First(&pendidikan, newKuis.Pendidikan_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Pendidikan ID", nil)
	}

	// Create Kuis using database function
	result, err := database.CreateKuis(newKuis.Title, newKuis.Description, newKuis.IsPrivate, newKuis.Kategori_id, newKuis.Tingkatan_id, newKuis.Kelas_id, newKuis.Pendidikan_id, user.ID)
	if err != nil {
		return handleError(c, err, "Failed to create quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz created successfully", result)
}

func UpdateKuis(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse request body
	newTask := new(models.Kuis)
	err := c.BodyParser(newTask)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdateKuis(newTask.Title, newTask.Description, newTask.IsPrivate, newTask.Kategori_id, newTask.Tingkatan_id, newTask.Kelas_id, newTask.Pendidikan_id, id)
	if err != nil {
		return handleError(c, err, "Failed to update quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz updated successfully", result)
}

func DeleteKuis(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err := database.DeleteKuis(id)
	if err != nil {
		return handleError(c, err, "Failed to delete quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz deleted successfully", nil)
}
func FilterKuis(c *fiber.Ctx) error {
	// Ambil parameter dari query string
	kategoriID := c.Query("kategori_id")     // Misalnya ?kategori_id=1
	tingkatanID := c.Query("tingkatan_id")   // Misalnya ?tingkatan_id=1
	pendidikanID := c.Query("pendidikan_id") // Misalnya ?pendidikan_id=1

	// Membuat query untuk filter
	var kuis []models.Kuis
	query := database.DB.Model(&models.Kuis{})

	// Jika kategori_id disediakan, filter berdasarkan kategori_id
	if kategoriID != "" {
		query = query.Where("kategori_id = ?", kategoriID)
	}

	// Jika tingkatan_id disediakan, filter berdasarkan tingkatan_id
	if tingkatanID != "" {
		query = query.Where("tingkatan_id = ?", tingkatanID)
	}

	// Jika pendidikan_id disediakan, filter berdasarkan pendidikan_id
	if pendidikanID != "" {
		query = query.Where("pendidikan_id = ?", pendidikanID)
	}

	// Menjalankan query untuk mendapatkan kuis yang sesuai dengan filter
	err := query.Find(&kuis).Error
	if err != nil {
		return sendResponse(c, fiber.StatusInternalServerError, false, "Failed to fetch quizzes", nil)
	}

	// Mengembalikan daftar kuis yang telah difilter
	return sendResponse(c, fiber.StatusOK, true, "Filtered quizzes retrieved successfully", kuis)
}
