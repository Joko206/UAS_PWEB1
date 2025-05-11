package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

// Fungsi untuk user bergabung ke kelas
func JoinKelas(c *fiber.Ctx) error {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Parse data dari body (user_id dan kelas_id)
	var requestData struct {
		User_id  uint `json:"user_id"`  // ID pengguna
		Kelas_id uint `json:"kelas_id"` // ID kelas
	}

	// Parse body request untuk mendapatkan user_id dan kelas_id
	err = c.BodyParser(&requestData)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Periksa apakah user sudah terdaftar di kelas ini
	var existingRecord models.Kelas_Pengguna
	err = db.Where("users_id = ? AND kelas_id = ?", requestData.User_id, requestData.Kelas_id).First(&existingRecord).Error
	if err == nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "User already joined this class",
		})
	}

	// Simpan data user yang bergabung ke dalam kelas
	newRecord := models.Kelas_Pengguna{
		Users_id: requestData.User_id,
		Kelas_id: requestData.Kelas_id,
	}

	if err := db.Create(&newRecord).Error; err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Failed to join the class",
		})
	}

	// Return success response
	return c.Status(200).JSON(&fiber.Map{
		"data":    newRecord,
		"success": true,
		"message": "User joined the class successfully",
	})
}
