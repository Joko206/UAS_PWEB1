package database

import (
	"encoding/json"
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateSoal(question string, option json.RawMessage, correct_answer string, kuis_id int) (models.Soal, error) {
	// Create a new Kategori_Soal instance
	var newKategori = models.Soal{Question: question, Options: option, Correct_answer: correct_answer, Kuis_id: kuis_id}

	// Open a database connection (or reuse the global DB connection)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return newKategori, err
	}

	// Insert the new category into the database
	err = db.Create(&newKategori).Error
	if err != nil {
		log.Fatal("Error inserting data into kategori_soal:", err)
		return newKategori, err
	}

	// Return the newly created category
	return newKategori, nil
}
func GetSoal() ([]models.Soal, error) {
	var getKategori []models.Soal

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return getKategori, err
	}

	db.Find(&getKategori)

	return getKategori, nil
}
func DeletSoal(id string) error {
	var deleteKategori models.Soal

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&deleteKategori)
	return nil

}
func UpdateSoal(question string, option json.RawMessage, correct_answer string, kuis_id int, id string) (models.Soal, error) {
	var newTask = models.Soal{Question: question, Options: option, Correct_answer: correct_answer, Kuis_id: kuis_id}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Soal{Question: newTask.Question, Options: newTask.Options, Correct_answer: newTask.Correct_answer, Kuis_id: newTask.Kuis_id})
	return newTask, nil
}
