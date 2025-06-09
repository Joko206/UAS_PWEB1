package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func JoinKelas(c *fiber.Ctx) error {
	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	var requestData struct {
		User_id  uint `json:"user_id"`
		Kelas_id uint `json:"kelas_id"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Cek apakah user dengan User_id ada
	var user models.Users
	if err := db.First(&user, requestData.User_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "User does not exist", nil)
	}

	// Cek apakah kelas dengan Kelas_id ada
	var kelas models.Kelas
	if err := db.First(&kelas, requestData.Kelas_id).Error; err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Class does not exist", nil)
	}

	// Cek apakah user sudah tergabung dengan kelas
	var existingRecord models.Kelas_Pengguna
	if err := db.Where("users_id = ? AND kelas_id = ?", requestData.User_id, requestData.Kelas_id).First(&existingRecord).Error; err == nil {
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
	// Ambil user yang sedang login
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Ambil semua kelas yang diikuti oleh user dengan preload untuk efisiensi
	var kelasPengguna []models.Kelas_Pengguna
	if err := db.Preload("Kelas").Where("users_id = ?", user.ID).Find(&kelasPengguna).Error; err != nil {
		return handleError(c, err, "Failed to get user classes")
	}

	// Extract kelas data dari relasi
	var kelasList []models.Kelas
	for _, kp := range kelasPengguna {
		kelasList = append(kelasList, kp.Kelas)
	}

	// Return response menggunakan helper function
	return sendResponse(c, fiber.StatusOK, true, "User classes retrieved successfully", kelasList)
}
