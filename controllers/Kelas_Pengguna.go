package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func JoinKelas(c *fiber.Ctx) error {

	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	var requestData struct {
		User_id  uint `json:"user_id"`
		Kelas_id uint `json:"kelas_id"`
	}

	err = c.BodyParser(&requestData)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	var existingRecord models.Kelas_Pengguna
	err = db.Where("users_id = ? AND kelas_id = ?", requestData.User_id, requestData.Kelas_id).First(&existingRecord).Error
	if err == nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User already joined this class", nil)
	}

	newRecord := models.Kelas_Pengguna{
		Users_id: requestData.User_id,
		Kelas_id: requestData.Kelas_id,
	}

	if err := db.Create(&newRecord).Error; err != nil {
		return handleError(c, err, "Failed to join the class")
	}

	return sendResponse(c, fiber.StatusOK, true, "User joined the class successfully", newRecord)
}
