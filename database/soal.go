package database

import (
	"encoding/json"
	"fmt"
	"github.com/Joko206/UAS_PWEB1/models"
)

func CreateSoal(question string, option json.RawMessage, correct_answer string, kuis_id uint) (models.Soal, error) {
	var newSoal = models.Soal{
		Question:       question,
		Options:        option,
		Correct_answer: correct_answer,
		Kuis_id:        kuis_id,
	}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newSoal, err
	}

	// Insert the new Soal into the database
	if err := db.Create(&newSoal).Error; err != nil {
		return newSoal, fmt.Errorf("failed to insert data into soal: %w", err)
	}

	return newSoal, nil
}

// GetSoal retrieves all Soal from the database
func GetSoal() ([]models.Soal, error) {
	var soalList []models.Soal

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return soalList, err
	}

	// Retrieve all Soal
	if err := db.Find(&soalList).Error; err != nil {
		return soalList, fmt.Errorf("failed to retrieve soal: %w", err)
	}

	return soalList, nil
}

// DeleteSoal deletes a Soal by its ID
func DeleteSoal(id string) error {
	var soal models.Soal

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the Soal by ID
	if err := db.Where("ID = ?", id).Delete(&soal).Error; err != nil {
		return fmt.Errorf("failed to delete soal: %w", err)
	}

	return nil
}

// UpdateSoal updates an existing Soal in the database
func UpdateSoal(question string, option json.RawMessage, correct_answer string, kuis_id uint, id string) (models.Soal, error) {
	var updatedSoal = models.Soal{
		Question:       question,
		Options:        option,
		Correct_answer: correct_answer,
		Kuis_id:        kuis_id,
	}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedSoal, err
	}

	// Update the Soal details
	if err := db.Where("ID = ?", id).Updates(&updatedSoal).Error; err != nil {
		return updatedSoal, fmt.Errorf("failed to update soal: %w", err)
	}

	return updatedSoal, nil
}
