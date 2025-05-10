package database

import (
	"log"

	"belajar-via-dev.to/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateTingkatan(name string, description string) (models.Tingkatan, error) {
	// Create a new Kategori_Soal instance
	var newTask = models.Tingkatan{Name: name, Description: description}

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
func GetTingkatan() ([]models.Tingkatan, error) {
	var newTask []models.Tingkatan

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Find(&newTask)

	return newTask, nil
}
func DeleteTingkatan(id string) error {
	var newTask models.Tingkatan

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&newTask)
	return nil

}
func UpdateTingkatan(name string, description string, id string) (models.Tingkatan, error) {
	var newTask = models.Tingkatan{Name: name, Description: description}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Tingkatan{Name: newTask.Name, Description: newTask.Description})
	return newTask, nil
}
