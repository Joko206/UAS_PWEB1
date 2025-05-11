package database

import (
	"fmt"
	"github.com/Joko206/UAS_PWEB1/models"
)

// CreatePendidikan creates a new Pendidikan in the database
func CreatePendidikan(name string, description string) (models.Pendidikan, error) {
	var newPendidikan = models.Pendidikan{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newPendidikan, err
	}

	// Insert the new Pendidikan into the database
	if err := db.Create(&newPendidikan).Error; err != nil {
		return newPendidikan, fmt.Errorf("failed to insert data into pendidikan: %w", err)
	}

	return newPendidikan, nil
}

// GetPendidikan retrieves all Pendidikan from the database
func GetPendidikan() ([]models.Pendidikan, error) {
	var pendidikanList []models.Pendidikan

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return pendidikanList, err
	}

	// Retrieve all Pendidikan
	if err := db.Find(&pendidikanList).Error; err != nil {
		return pendidikanList, fmt.Errorf("failed to retrieve pendidikan: %w", err)
	}

	return pendidikanList, nil
}

// DeletePendidikan deletes a Pendidikan by its ID
func DeletePendidikan(id string) error {
	var pendidikan models.Pendidikan

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the Pendidikan by ID
	if err := db.Where("ID = ?", id).Delete(&pendidikan).Error; err != nil {
		return fmt.Errorf("failed to delete pendidikan: %w", err)
	}

	return nil
}

// UpdatePendidikan updates an existing Pendidikan in the database
func UpdatePendidikan(name string, description string, id string) (models.Pendidikan, error) {
	var updatedPendidikan = models.Pendidikan{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedPendidikan, err
	}

	// Update the Pendidikan details
	if err := db.Where("ID = ?", id).Updates(&updatedPendidikan).Error; err != nil {
		return updatedPendidikan, fmt.Errorf("failed to update pendidikan: %w", err)
	}

	return updatedPendidikan, nil
}
