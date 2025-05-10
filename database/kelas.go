package database

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateKelas(name string, description string) (models.Kelas, error) {
	// Create a new Kategori_Soal instance
	var newTask = models.Kelas{Name: name, Description: description}

	// Open a database connection (or reuse the global DB connection)
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
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
func GetKelas() ([]models.Kelas, error) {
	var newTask []models.Kelas

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Find(&newTask)

	return newTask, nil
}
func DeleteKelas(id string) error {
	var newTask models.Kelas

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&newTask)
	return nil

}
func UpdateKelas(name string, description string, id string) (models.Kelas, error) {
	var newTask = models.Kelas{Name: name, Description: description}

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Kelas{Name: newTask.Name, Description: newTask.Description})
	return newTask, nil
}
