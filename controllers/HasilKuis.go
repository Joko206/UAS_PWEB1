package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func SubmitJawaban(c *fiber.Ctx) error {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Parse data dari body (jawaban yang diberikan oleh user)
	var userAnswers []models.SoalAnswer
	err = c.BodyParser(&userAnswers)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	// Simpan jawaban pengguna ke dalam SoalAnswer
	if err := db.Create(&userAnswers).Error; err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Failed to save answers",
		})
	}

	// Ambil kuis_id dari soal yang pertama
	soalID := userAnswers[0].Soal_id
	var soal models.Soal
	if err := db.First(&soal, soalID).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Invalid Soal ID",
		})
	}

	// Ambil kuis_id dari soal yang terkait
	kuisID := soal.Kuis_id

	// Dapatkan soal-soal yang terkait dengan kuis ini
	var soalList []models.Soal
	if err := db.Where("kuis_id = ?", kuisID).Find(&soalList).Error; err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	// Hitung skor dan jumlah jawaban yang benar
	var correctAnswers uint
	for _, answer := range userAnswers {
		for _, soal := range soalList {
			// Perbandingkan soal_id dan jawaban dengan soal yang benar
			if answer.Soal_id == soal.ID && answer.Answer == soal.Correct_answer {
				correctAnswers++
			}
		}
	}

	// Hitung skor
	score := correctAnswers * 10 // Misalnya 10 poin untuk setiap jawaban yang benar

	// Simpan hasil kuis ke tabel Hasil_Kuis
	result := models.Hasil_Kuis{
		Users_id:       userAnswers[0].User_id,
		Kuis_id:        kuisID,
		Score:          score,
		Correct_Answer: correctAnswers,
	}

	// Pastikan tidak ada duplikasi hasil untuk user_id dan kuis_id yang sama
	var existingResult models.Hasil_Kuis
	if err := db.Where("users_id = ? AND kuis_id = ?", userAnswers[0].User_id, kuisID).First(&existingResult).Error; err == nil {
		// Jika hasil sudah ada, update hasil yang lama
		existingResult.Score = score
		existingResult.Correct_Answer = correctAnswers
		if err := db.Save(&existingResult).Error; err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"data":    nil,
				"success": false,
				"message": "Failed to update result",
			})
		}
	} else {
		// Jika hasil belum ada, simpan hasil baru
		if err := db.Create(&result).Error; err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"data":    nil,
				"success": false,
				"message": "Failed to save result",
			})
		}
	}

	// Return hasilnya
	return c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Kuis submitted successfully",
	})
}

// Fungsi untuk mendapatkan hasil kuis berdasarkan user_id dan kuis_id
func GetHasilKuis(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	kuisID := c.Params("kuis_id")

	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Failed to connect to the database",
		})
	}

	// Cari hasil kuis berdasarkan user_id dan kuis_id
	var hasilKuis models.Hasil_Kuis
	if err := db.Where("users_id = ? AND kuis_id = ?", userID, kuisID).First(&hasilKuis).Error; err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Result not found",
		})
	}

	// Kembalikan hasil kuis
	return c.Status(200).JSON(&fiber.Map{
		"data":    hasilKuis,
		"success": true,
		"message": "Hasil kuis ditemukan",
	})

}
