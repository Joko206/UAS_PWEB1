package database

import (
	"fmt"

	"github.com/Joko206/UAS_PWEB1/models"
)

// CreateKuis creates a new Kuis in the database
func CreateKuis(title string, description string, kategori uint, tingkatan uint, kelas uint, pendidikan uint) (models.Kuis, error) {
	var newKuis = models.Kuis{
		Title:        title,
		Description:  description,
		KategoriID:   kategori,
		TingkatanID:  tingkatan,
		KelasID:      kelas,
		PendidikanID: pendidikan,
	}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return newKuis, err
	}

	// Validate Kategori, Tingkatan, and Kelas
	var kategoriObj models.Kategori_Soal
	if err := db.First(&kategoriObj, kategori).Error; err != nil {
		return newKuis, fmt.Errorf("Invalid Kategori ID")
	}

	var tingkatanObj models.Tingkatan
	if err := db.First(&tingkatanObj, tingkatan).Error; err != nil {
		return newKuis, fmt.Errorf("Invalid Tingkatan ID")
	}

	var kelasObj models.Kelas
	if err := db.First(&kelasObj, kelas).Error; err != nil {
		return newKuis, fmt.Errorf("Invalid Kelas ID")
	}

	var pendidikanObj models.Kelas
	if err := db.First(&pendidikanObj, pendidikan).Error; err != nil {
		return newKuis, fmt.Errorf("Invalid Kelas ID")
	}

	// Insert the new Kuis into the database
	if err := db.Create(&newKuis).Error; err != nil {
		return newKuis, fmt.Errorf("failed to insert data into kuis: %w", err)
	}

	return newKuis, nil
}

// GetKuis retrieves all Kuis from the database with related Kategori, Tingkatan, and Kelas
func GetKuis() ([]models.Kuis, error) {
	var kuisList []models.Kuis

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return kuisList, err
	}

	// Preload related models (Kategori, Tingkatan, Kelas)
	if err := db.Preload("Kategori").Preload("Tingkatan").Preload("Kelas").Find(&kuisList).Error; err != nil {
		return kuisList, fmt.Errorf("failed to retrieve kuis: %w", err)
	}

	return kuisList, nil
}

// DeleteKuis deletes a Kuis by its ID
func DeleteKuis(id string) error {
	var kuis models.Kuis

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Delete the Kuis by ID
	if err := db.Where("ID = ?", id).Delete(&kuis).Error; err != nil {
		return fmt.Errorf("failed to delete kuis: %w", err)
	}

	return nil
}

// UpdateKuis updates an existing Kuis in the database
func UpdateKuis(title string, description string, kategori uint, tingkatan uint, kelas uint, id string) (models.Kuis, error) {
	var updatedKuis = models.Kuis{
		Title:       title,
		Description: description,
		KategoriID:  kategori,
		TingkatanID: tingkatan,
		KelasID:     kelas,
	}

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return updatedKuis, err
	}

	// Update the Kuis details
	if err := db.Where("ID = ?", id).Updates(&updatedKuis).Error; err != nil {
		return updatedKuis, fmt.Errorf("failed to update kuis: %w", err)
	}

	return updatedKuis, nil
}
