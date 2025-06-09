package database

import (
	"fmt"

	"github.com/Joko206/UAS_PWEB1/models"
)

// CreateKategori creates a new Kategori_Soal in the database
func CreateKategori(name string, description string) (models.Kategori_Soal, error) {
	var newKategori = models.Kategori_Soal{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newKategori, err
	}

	// Insert the new category into the database
	if err := db.Create(&newKategori).Error; err != nil {
		return newKategori, fmt.Errorf("failed to insert data into kategori_soal: %w", err)
	}

	return newKategori, nil
}

// GetKategori retrieves all Kategori_Soal from the database
func GetKategori() ([]models.Kategori_Soal, error) {
	var getKategori []models.Kategori_Soal

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return getKategori, err
	}

	// Retrieve all categories
	if err := db.Find(&getKategori).Error; err != nil {
		return getKategori, fmt.Errorf("failed to retrieve categories: %w", err)
	}

	return getKategori, nil
}

// DeleteKategori deletes a Kategori_Soal by its ID
func DeleteKategori(id string) error {
	var deleteKategori models.Kategori_Soal

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the category by ID
	if err := db.Where("ID = ?", id).Delete(&deleteKategori).Error; err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

// UpdateKategori updates an existing Kategori_Soal in the database
func UpdateKategori(name string, description string, id string) (models.Kategori_Soal, error) {
	var updatedKategori = models.Kategori_Soal{Name: name, Description: description}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedKategori, err
	}

	// Update the category details
	if err := db.Where("ID = ?", id).Updates(&updatedKategori).Error; err != nil {
		return updatedKategori, fmt.Errorf("failed to update category: %w", err)
	}

	return updatedKategori, nil
}
