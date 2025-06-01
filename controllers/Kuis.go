package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetKuis(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetKuis()
	if err != nil {
		return handleError(c, err, "Failed to retrieve quizzes")
	}

	return sendResponse(c, fiber.StatusOK, true, "All quizzes retrieved successfully", result)
}

func AddKuis(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Use existing database connection
	db := database.MustGetDB()

	// Parse request body
	newKuis := new(models.Kuis)
	err = c.BodyParser(newKuis)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Validate Kategori
	var kategori models.Kategori_Soal
	if err := db.First(&kategori, newKuis.KategoriID).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Kategori ID", nil)
	}

	// Validate Tingkatan
	var tingkatan models.Tingkatan
	if err := db.First(&tingkatan, newKuis.TingkatanID).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Tingkatan ID", nil)
	}

	// Validate Kelas
	var kelas models.Kelas
	if err := db.First(&kelas, newKuis.KelasID).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Kelas ID", nil)
	}

	var pendidikan models.Pendidikan
	if err := db.First(&pendidikan, newKuis.PendidikanID).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid Pendidikan ID", nil)
	}

	// Create Kuis
	result, err := database.CreateKuis(newKuis.Title, newKuis.Description, newKuis.KategoriID, newKuis.TingkatanID, newKuis.KelasID, newKuis.PendidikanID)
	if err != nil {
		return handleError(c, err, "Failed to create quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz created successfully", result)
}

func UpdateKuis(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	// Parse request body
	newTask := new(models.Kuis)
	err = c.BodyParser(newTask)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdateKuis(newTask.Title, newTask.Description, newTask.KategoriID, newTask.TingkatanID, newTask.KelasID, id)
	if err != nil {
		return handleError(c, err, "Failed to update quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz updated successfully", result)
}

func DeleteKuis(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err = database.DeleteKuis(id)
	if err != nil {
		return handleError(c, err, "Failed to delete quiz")
	}

	return sendResponse(c, fiber.StatusOK, true, "Quiz deleted successfully", nil)
}
func FilterKuis(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Ambil parameter dari query string
	kategoriID := c.Query("kategori_id")   // Misalnya ?kategori_id=1
	tingkatanID := c.Query("tingkatan_id") // Misalnya ?tingkatan_id=1
	kelasID := c.Query("kelas_id")
	pendidikanID := c.Query("pendidikan_id") // Misalnya ?pendidikan_id=1

	// Membuat query untuk filter
	var kuis []models.Kuis
	db := database.MustGetDB()
	query := db.Model(&models.Kuis{})

	// Jika kategori_id disediakan, filter berdasarkan kategori_id
	if kategoriID != "" {
		query = query.Where("kategori_id = ?", kategoriID)
	}

	// Jika tingkatan_id disediakan, filter berdasarkan tingkatan_id
	if tingkatanID != "" {
		query = query.Where("tingkatan_id = ?", tingkatanID)
	}

	if kelasID != "" {
		query = query.Where("kelas_id = ?", kelasID)
	}

	// Jika pendidikan_id disediakan, filter berdasarkan pendidikan_id
	if pendidikanID != "" {
		query = query.Where("pendidikan_id = ?", pendidikanID)
	}

	// Menjalankan query untuk mendapatkan kuis yang sesuai dengan filter
	err = query.Find(&kuis).Error
	if err != nil {
		return sendResponse(c, fiber.StatusInternalServerError, false, "Failed to fetch quizzes", nil)
	}

	// Mengembalikan daftar kuis yang telah difilter
	return sendResponse(c, fiber.StatusOK, true, "Filtered quizzes retrieved successfully", kuis)
}
