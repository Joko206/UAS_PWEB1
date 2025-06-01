package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func JoinKelas(c *fiber.Ctx) error {
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Use the existing database connection
	db := database.MustGetDB()

	var requestData struct {
		UserID  uint `json:"user_id"`
		KelasID uint `json:"kelas_id"`
	}

	err = c.BodyParser(&requestData)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Cek apakah user dengan UserID ada
	var user models.Users
	err = db.First(&user, requestData.UserID).Error
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User does not exist", nil)
	}

	// Cek apakah kelas dengan KelasID ada
	var kelas models.Kelas
	err = db.First(&kelas, requestData.KelasID).Error
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Class does not exist", nil)
	}

	// Cek apakah user sudah tergabung dengan kelas
	var existingRecord models.KelasPengguna
	err = db.Where("user_id = ? AND kelas_id = ?", requestData.UserID, requestData.KelasID).First(&existingRecord).Error
	if err == nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User already joined this class", nil)
	}

	// Menambahkan data ke kelas_penggunas
	newRecord := models.KelasPengguna{
		UserID:  requestData.UserID,
		KelasID: requestData.KelasID,
	}

	if err := db.Create(&newRecord).Error; err != nil {
		return handleError(c, err, "Failed to join the class")
	}

	return sendResponse(c, fiber.StatusOK, true, "User joined the class successfully", newRecord)
}

func GetKelasByUserID(c *fiber.Ctx) error {
	// Ambil user yang sedang login
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Use the existing database connection
	db := database.MustGetDB()

	// Ambil semua kelas yang diikuti oleh user berdasarkan UserID
	var kelasPengguna []models.KelasPengguna
	err = db.Where("user_id = ?", user.ID).Find(&kelasPengguna).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get classes",
		})
	}

	// Ambil data kelas terkait
	var kelasList []models.Kelas
	for _, kp := range kelasPengguna {
		var kelas models.Kelas
		err := db.Where("id = ?", kp.KelasID).First(&kelas).Error
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
