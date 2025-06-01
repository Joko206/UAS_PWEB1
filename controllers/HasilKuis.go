package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func SubmitJawaban(c *fiber.Ctx) error {
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Gunakan database.DB yang sudah terkonfigurasi
	db := database.MustGetDB()

	// Parse data dari body (jawaban yang diberikan oleh user)
	var userAnswers []models.SoalAnswer
	err = c.BodyParser(&userAnswers)
	if err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Validasi jika tidak ada jawaban yang diberikan
	if len(userAnswers) == 0 {
		return sendResponse(c, fiber.StatusBadRequest, false, "No answers provided", nil)
	}

	// Ambil kuis_id dari soal pertama
	soalID := userAnswers[0].SoalID
	var soal models.Soal
	if err := db.First(&soal, soalID).Error; err != nil {
		return handleError(c, err, "Invalid Soal ID")
	}

	// Ambil kuis_id dari soal yang terkait
	kuisID := soal.KuisID

	// Dapatkan soal-soal yang terkait dengan kuis ini
	var soalList []models.Soal
	if err := db.Where("kuis_id = ?", kuisID).Find(&soalList).Error; err != nil {
		return handleError(c, err, "Failed to fetch related questions")
	}

	// Hitung skor dan jumlah jawaban yang benar
	var correctAnswers uint
	for _, answer := range userAnswers {
		for _, soal := range soalList {
			if answer.SoalID == soal.ID && answer.Answer == soal.CorrectAnswer {
				correctAnswers++
			}
		}
	}

	// Hitung skor
	score := correctAnswers * 10 // Misalnya 10 poin untuk setiap jawaban yang benar

	// Simpan hasil kuis ke tabel HasilKuis
	result := models.HasilKuis{
		UserID:        userAnswers[0].UserID,
		KuisID:        kuisID,
		Score:         score,
		CorrectAnswer: correctAnswers,
	}

	// Cek apakah hasil sudah ada
	var existingResult models.HasilKuis
	if err := db.Where("user_id = ? AND kuis_id = ?", userAnswers[0].UserID, kuisID).First(&existingResult).Error; err == nil {
		// Jika sudah ada, update hasil
		existingResult.Score = score
		existingResult.CorrectAnswer = correctAnswers
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

func GetHasilKuis(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	kuisID := c.Params("kuis_id")

	// Gunakan database connection yang sudah ada
	db := database.MustGetDB()

	// Cari hasil kuis berdasarkan user_id dan kuis_id
	var hasilKuis models.HasilKuis
	if err := db.Where("user_id = ? AND kuis_id = ?", userID, kuisID).First(&hasilKuis).Error; err != nil {
		return sendResponse(c, fiber.StatusNotFound, false, "Result not found", nil)
	}

	// Kembalikan hasil kuis
	return sendResponse(c, fiber.StatusOK, true, "Hasil kuis ditemukan", hasilKuis)
}
