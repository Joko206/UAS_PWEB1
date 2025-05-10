package database

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePendidikan(name string, description string) (models.Pendidikan, error) {
	// Create a new Kategori_Soal instance
	var newPendidikan = models.Pendidikan{Name: name, Description: description}

	// Open a database connection (or reuse the global DB connection)
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return newPendidikan, err
	}

	// Insert the new category into the database
	err = db.Create(&newPendidikan).Error
	if err != nil {
		log.Fatal("Error inserting data into kategori_soal:", err)
		return newPendidikan, err
	}

	// Return the newly created category
	return newPendidikan, nil
}
func GetPendidikan() ([]models.Pendidikan, error) {
	var getKategori []models.Pendidikan

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return getKategori, err
	}

	db.Find(&getKategori)

	return getKategori, nil
}

func DeletePendidikan(id string) error {
	var deleteKategori models.Pendidikan

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&deleteKategori)
	return nil

}
func UpdatePendidikan(name string, description string, id string) (models.Pendidikan, error) {
	var newTask = models.Pendidikan{Name: name, Description: description}

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Kategori_Soal{Name: newTask.Name, Description: newTask.Description})
	return newTask, nil
}
