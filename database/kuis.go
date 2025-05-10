package database

import (
	"fmt"
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Fungsi untuk membuat Kuis
func CreateKuis(title string, description string, kategori uint, tingkatan uint, kelas uint) (models.Kuis, error) {
	var newTask = models.Kuis{Title: title, Description: description, Kategori_id: kategori, Tingkatan_id: tingkatan, Kelas_id: kelas}

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	// Pastikan ID kategori, tingkatan, dan kelas valid
	var kategoriObj models.Kategori_Soal
	if err := db.First(&kategoriObj, kategori).Error; err != nil {
		return newTask, fmt.Errorf("Invalid Kategori ID")
	}

	// Lakukan hal yang sama untuk Tingkatan dan Kelas
	var tingkatanObj models.Tingkatan
	if err := db.First(&tingkatanObj, tingkatan).Error; err != nil {
		return newTask, fmt.Errorf("Invalid Tingkatan ID")
	}

	var kelasObj models.Kelas
	if err := db.First(&kelasObj, kelas).Error; err != nil {
		return newTask, fmt.Errorf("Invalid Kelas ID")
	}

	// Setelah validasi, masukkan data
	err = db.Create(&newTask).Error
	if err != nil {
		log.Fatal("Error inserting data into kuis:", err)
		return newTask, err
	}

	return newTask, nil
}

func GetKuis() ([]models.Kuis, error) {
	var newTask []models.Kuis

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Find(&newTask)

	return newTask, nil
}
func DeleteKuis(id string) error {
	var newTask models.Kuis

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&newTask)
	return nil

}
func UpdateKuis(title string, description string, kategori uint, tingkatan uint, kelas uint, id string) (models.Kuis, error) {
	var newTask = models.Kuis{Title: title, Description: description, Kategori_id: kategori, Tingkatan_id: tingkatan, Kelas_id: kelas}

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Kuis{Title: newTask.Title, Description: newTask.Description, Kategori_id: newTask.Kategori_id, Tingkatan_id: newTask.Tingkatan_id, Kelas_id: newTask.Kelas_id})
	return newTask, nil

}
