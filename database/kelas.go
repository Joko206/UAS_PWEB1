package database

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Joko206/UAS_PWEB1/models"
)

// generateJoinCode generates a unique 6-character join code
func generateJoinCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range 6 {
		result.WriteByte(charset[r.Intn(len(charset))])
	}
	return result.String()
}

// CreateKelas creates a new Kelas in the database with join code
func CreateKelas(name string, description string, createdBy uint) (models.Kelas, error) {
	joinCode := generateJoinCode()

	var newKelas = models.Kelas{
		Name:        name,
		Description: description,
		JoinCode:    joinCode,
		CreatedBy:   createdBy,
	}

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

// GetKelasByJoinCode finds a class by its join code
func GetKelasByJoinCode(joinCode string) (models.Kelas, error) {
	var kelas models.Kelas

	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return kelas, err
	}

	// Find class by join code
	if err := db.Where("join_code = ?", joinCode).First(&kelas).Error; err != nil {
		return kelas, fmt.Errorf("class with join code %s not found", joinCode)
	}

	return kelas, nil
}

// JoinKelasByCode allows a user to join a class using join code
func JoinKelasByCode(userID uint, joinCode string) error {
	// Get DB connection
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	// Find class by join code
	var kelas models.Kelas
	if err := db.Where("join_code = ?", joinCode).First(&kelas).Error; err != nil {
		return fmt.Errorf("invalid join code")
	}

	// Check if user already joined this class
	var existingRecord models.Kelas_Pengguna
	err = db.Where("users_id = ? AND kelas_id = ?", userID, kelas.ID).First(&existingRecord).Error
	if err == nil {
		return fmt.Errorf("user already joined this class")
	}

	// Add user to class
	newRecord := models.Kelas_Pengguna{
		Users_id: userID,
		Kelas_id: kelas.ID,
	}

	if err := db.Create(&newRecord).Error; err != nil {
		return fmt.Errorf("failed to join class: %w", err)
	}

	return nil
}
