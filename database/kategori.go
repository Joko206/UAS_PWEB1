package database

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateKategori(name string, description string) (models.Kategori_Soal, error) {
	// Create a new Kategori_Soal instance
	var newKategori = models.Kategori_Soal{Name: name, Description: description}

	// Open a database connection (or reuse the global DB connection)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return newKategori, err
	}

	// Insert the new category into the database
	err = db.Create(&newKategori).Error
	if err != nil {
		log.Fatal("Error inserting data into kategori_soal:", err)
		return newKategori, err
	}

	// Return the newly created category
	return newKategori, nil
}
func GetallTasks() ([]models.Kategori_Soal, error) {
	var getKategori []models.Kategori_Soal

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return getKategori, err
	}

	db.Find(&getKategori)

	return getKategori, nil
}
func DeleteKateggori(id string) error {
	var deleteKategori models.Kategori_Soal

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("ID = ?", id).Delete(&deleteKategori)
	return nil

}
func UpdateKategori(name string, description string, id string) (models.Kategori_Soal, error) {
	var newTask = models.Kategori_Soal{Name: name, Description: description}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}

	db.Where("ID = ?", id).Updates(&models.Kategori_Soal{Name: newTask.Name, Description: newTask.Description})
	return newTask, nil
}
