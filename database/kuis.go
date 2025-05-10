package database

import (
	"log"

	"belajar-via-dev.to/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateKuis(title string, description string, kategori_id int, tingkatan_id int, kelas_id int) (models.Kuis, error) {
	// Create a new Kategori_Soal instance
	var newTask = models.Kuis{Title: title, Description: description, Kategori_id: kategori_id, Tingkatan_id: tingkatan_id, Kelas_id: kelas_id}

	// Open a database connection (or reuse the global DB connection)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return newTask, err
	}

	// Insert the new category into the database
	err = db.Create(&newTask).Error
	if err != nil {
		log.Fatal("Error inserting data into kategori_soal:", err)
		return newTask, err
	}

	// Return the newly created category
	return newTask, nil
}
func GetKuis() ([]models.Kuis, error) {
	var newTask []models.Kuis

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Find(&newTask)

	return newTask, nil
}
func DeleteKuis(id string) error {
	var newTask models.Kuis

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&newTask)
	return nil

}
func UpdateKuis(title string, description string, kategori_id int, tingkatan_id int, kelas_id int, id string) (models.Kuis, error) {
	var newTask = models.Kuis{Title: title, Description: description, Kategori_id: kategori_id, Tingkatan_id: tingkatan_id, Kelas_id: kelas_id}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Kuis{Title: newTask.Title, Description: newTask.Description, Kategori_id: newTask.Kategori_id, Tingkatan_id: newTask.Tingkatan_id, Kelas_id: newTask.Kelas_id})
	return newTask, nil

}
