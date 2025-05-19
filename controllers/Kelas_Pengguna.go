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

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
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

	// Cek apakah user dengan User_id ada
	var user models.Users
	err = db.First(&user, requestData.User_id).Error
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User does not exist", nil)
	}

	// Cek apakah kelas dengan Kelas_id ada
	var kelas models.Kelas
	err = db.First(&kelas, requestData.Kelas_id).Error
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Class does not exist", nil)
	}

	// Cek apakah user sudah tergabung dengan kelas
	var existingRecord models.Kelas_Pengguna
	err = db.Where("users_id = ? AND kelas_id = ?", requestData.User_id, requestData.Kelas_id).First(&existingRecord).Error
	if err == nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User already joined this class", nil)
	}

	// Menambahkan data ke kelas_penggunas
	newRecord := models.Kelas_Pengguna{
		Users_id: requestData.User_id,
		Kelas_id: requestData.Kelas_id,
	}

	if err := db.Create(&newRecord).Error; err != nil {
		return handleError(c, err, "Failed to join the class")
	}

	return sendResponse(c, fiber.StatusOK, true, "User joined the class successfully", newRecord)
}

func GetKelasByUserID(c *fiber.Ctx) error {

	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Ambil user yang sedang login
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Ambil semua kelas yang diikuti oleh user berdasarkan Users_id
	var kelasPengguna []models.Kelas_Pengguna
	err = database.DB.Where("users_id = ?", user.ID).Find(&kelasPengguna).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get classes",
		})
	}

	// Ambil data kelas terkait
	var kelasList []models.Kelas
	for _, kp := range kelasPengguna {
		var kelas models.Kelas
		err := database.DB.Where("id = ?", kp.Kelas_id).First(&kelas).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get class data",
			})
		}
		kelasList = append(kelasList, kelas)
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    kelasList,
	})
}
