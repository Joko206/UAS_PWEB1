package database

import (
	"fmt"
	"github.com/Joko206/UAS_PWEB1/models"
)

// CreateTingkatan creates a new Tingkatan in the database
func CreateTingkatan(name string, description string) (models.Tingkatan, error) {
	var newTingkatan = models.Tingkatan{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newTingkatan, err
	}

	// Insert the new Tingkatan into the database
	if err := db.Create(&newTingkatan).Error; err != nil {
		return newTingkatan, fmt.Errorf("failed to insert data into tingkatan: %w", err)
	}

	return newTingkatan, nil
}

// GetTingkatan retrieves all Tingkatan from the database
func GetTingkatan() ([]models.Tingkatan, error) {
	var tingkatanList []models.Tingkatan

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return tingkatanList, err
	}

	// Retrieve all Tingkatan
	if err := db.Find(&tingkatanList).Error; err != nil {
		return tingkatanList, fmt.Errorf("failed to retrieve tingkatan: %w", err)
	}

	return tingkatanList, nil
}

// DeleteTingkatan deletes a Tingkatan by its ID
func DeleteTingkatan(id string) error {
	var tingkatan models.Tingkatan

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the Tingkatan by ID
	if err := db.Where("ID = ?", id).Delete(&tingkatan).Error; err != nil {
		return fmt.Errorf("failed to delete tingkatan: %w", err)
	}

	return nil
}

// UpdateTingkatan updates an existing Tingkatan in the database
func UpdateTingkatan(name string, description string, id string) (models.Tingkatan, error) {
	var updatedTingkatan = models.Tingkatan{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedTingkatan, err
	}

	// Update the Tingkatan details
	if err := db.Where("ID = ?", id).Updates(&updatedTingkatan).Error; err != nil {
		return updatedTingkatan, fmt.Errorf("failed to update tingkatan: %w", err)
	}

	return updatedTingkatan, nil
}
