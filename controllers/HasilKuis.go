package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func SubmitJawaban(c *fiber.Ctx) error {
	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Parse data dari body (jawaban yang diberikan oleh user)
	var userAnswers []models.SoalAnswer
	if err := c.BodyParser(&userAnswers); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Simpan jawaban pengguna ke dalam SoalAnswer
	if err := db.Create(&userAnswers).Error; err != nil {
		return handleError(c, err, "Failed to save answers")
	}

	// Ambil soal terkait untuk mendapatkan kuis_id
	soalID := userAnswers[0].Soal_id
	var soal models.Soal
	if err := db.First(&soal, soalID).Error; err != nil {
		return handleError(c, err, "Invalid Soal ID")
	}

	// Ambil kuis_id dari soal yang terkait
	kuisID := soal.Kuis_id

	// Dapatkan soal-soal yang terkait dengan kuis ini
	var soalList []models.Soal
	if err := db.Where("kuis_id = ?", kuisID).Find(&soalList).Error; err != nil {
		return handleError(c, err, "Failed to fetch related questions")
	}

	// Hitung skor dan jumlah jawaban yang benar
	var correctAnswers uint
	for _, answer := range userAnswers {
		for _, soal := range soalList {
			if answer.Soal_id == soal.ID && answer.Answer == soal.Correct_answer {
				correctAnswers++
			}
		}
	}

	// Hitung skor sebagai persentase (0-100)
	var score uint
	if len(soalList) > 0 {
		score = uint((float64(correctAnswers) / float64(len(soalList))) * 100)
	} else {
		score = 0
	}

	// Simpan hasil kuis ke tabel Hasil_Kuis
	result := models.Hasil_Kuis{
		Users_id:       userAnswers[0].User_id,
		Kuis_id:        kuisID,
		Score:          score,
		Correct_Answer: correctAnswers,
	}

	// Cek apakah hasil sudah ada
	var existingResult models.Hasil_Kuis
	if err := db.Where("users_id = ? AND kuis_id = ?", userAnswers[0].User_id, kuisID).First(&existingResult).Error; err == nil {
		// Jika sudah ada, update hasil
		existingResult.Score = score
		existingResult.Correct_Answer = correctAnswers
		if err := db.Save(&existingResult).Error; err != nil {
			return handleError(c, err, "Failed to update result")
		}
	} else {
		// Simpan hasil baru
		if err := db.Create(&result).Error; err != nil {
			return handleError(c, err, "Failed to save result")
		}
	}

	// Kembalikan hasil
	return sendResponse(c, fiber.StatusOK, true, "Kuis submitted successfully", result)
}

// GetHasilKuis - Get specific quiz result by user_id and kuis_id (kept for backward compatibility)
func GetHasilKuis(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	kuisID := c.Params("kuis_id")

	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Cari hasil kuis berdasarkan user_id dan kuis_id
	var hasilKuis models.Hasil_Kuis
	if err := db.Where("users_id = ? AND kuis_id = ?", userID, kuisID).First(&hasilKuis).Error; err != nil {
		return sendResponse(c, fiber.StatusNotFound, false, "Result not found", nil)
	}

	// Kembalikan hasil kuis
	return sendResponse(c, fiber.StatusOK, true, "Hasil kuis ditemukan", hasilKuis)
}

// GetAllHasilKuisByUser - Get all quiz results for the authenticated user (OPTIMIZED)
func GetAllHasilKuisByUser(c *fiber.Ctx) error {
	// Authenticate user
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Get all quiz results for the user with related quiz information in one query
	var hasilKuisList []models.Hasil_Kuis
	if err := db.Preload("Kuis").Where("users_id = ?", user.ID).Find(&hasilKuisList).Error; err != nil {
		return handleError(c, err, "Failed to fetch quiz results")
	}

	// Return all results
	return sendResponse(c, fiber.StatusOK, true, "All quiz results retrieved successfully", hasilKuisList)
}

// GetHasilKuisByUserID - Get all quiz results for a specific user (Admin/Teacher only)
func GetHasilKuisByUserID(c *fiber.Ctx) error {
	// Authenticate user
	authUser, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Check if user is admin or teacher
	if authUser.Role != "admin" && authUser.Role != "teacher" {
		return sendResponse(c, fiber.StatusForbidden, false, "Access denied. Admin or Teacher only.", nil)
	}

	userID := c.Params("user_id")
	if userID == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "User ID is required", nil)
	}

	// Get database connection (reuse global connection)
	db, err := database.GetDBConnection()
	if err != nil {
		return handleError(c, err, "Failed to get database connection")
	}

	// Get all quiz results for the specified user with related quiz information
	var hasilKuisList []models.Hasil_Kuis
	if err := db.Preload("Kuis").Where("users_id = ?", userID).Find(&hasilKuisList).Error; err != nil {
		return handleError(c, err, "Failed to fetch quiz results")
	}

	// Return all results
	return sendResponse(c, fiber.StatusOK, true, "Quiz results retrieved successfully", hasilKuisList)
}
