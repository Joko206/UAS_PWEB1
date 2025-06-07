package database

import (
	"fmt"
	"github.com/Joko206/UAS_PWEB1/models"
)

// CreateKelas creates a new Kelas in the database
func CreateKelas(name string, description string) (models.Kelas, error) {
	var newKelas = models.Kelas{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newKelas, err
	}

	// Insert the new class into the database
	if err := db.Create(&newKelas).Error; err != nil {
		return newKelas, fmt.Errorf("failed to insert data into kelas: %w", err)
	}

	return newKelas, nil
}

// GetKelas retrieves all Kelas from the database
func GetKelas() ([]models.Kelas, error) {
	var kelasList []models.Kelas

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return kelasList, err
	}

	// Retrieve all classes
	if err := db.Find(&kelasList).Error; err != nil {
		return kelasList, fmt.Errorf("failed to retrieve classes: %w", err)
	}

	return kelasList, nil
}

// DeleteKelas deletes a Kelas by its ID
func DeleteKelas(id string) error {
	var kelas models.Kelas

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the class by ID
	if err := db.Where("ID = ?", id).Delete(&kelas).Error; err != nil {
		return fmt.Errorf("failed to delete class: %w", err)
	}

	return nil
}

// UpdateKelas updates an existing Kelas in the database
func UpdateKelas(name string, description string, id string) (models.Kelas, error) {
	var updatedKelas = models.Kelas{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedKelas, err
	}

	// Update the class details
	if err := db.Where("ID = ?", id).Updates(&updatedKelas).Error; err != nil {
		return updatedKelas, fmt.Errorf("failed to update class: %w", err)
	}

	return updatedKelas, nil
}
