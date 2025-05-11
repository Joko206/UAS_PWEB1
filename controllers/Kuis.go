package controllers

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetKuis(c *fiber.Ctx) error {

	result, err := database.GetKuis()
	if err != nil {
		return handleError(c, err, "Failed to retrieve quizzes")
	}

	return sendResponse(c, fiber.StatusOK, true, "All quizzes retrieved successfully", result)
}

func AddKuis(c *fiber.Ctx) error {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	_, err = Authenticate(c)
	if err != nil {
		return err
	}

	// Parse request body
	newKuis := new(models.Kuis)
	err = c.BodyParser(newKuis)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
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

	// Create Kuis
	result, err := database.CreateKuis(newKuis.Title, newKuis.Description, newKuis.Kategori_id, newKuis.Tingkatan_id, newKuis.Kelas_id)
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

	result, err := database.UpdateKuis(newTask.Title, newTask.Description, newTask.Kategori_id, newTask.Tingkatan_id, newTask.Kelas_id, id)
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
