package controllers

import (
	"log"

	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetKuis(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetKuis()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "All Tasks",
	})
}

// Fungsi untuk menambahkan Kuis
func AddKuis(c *fiber.Ctx) error {

	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Authenticate the user using the JWT token
	_, err = Authenticate(c)
	if err != nil {
		return err
	}

	newKuis := new(models.Kuis)
	err = c.BodyParser(newKuis)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	// Cek apakah kategori_id ada di tabel kategori_soals
	var kategori models.Kategori_Soal
	if err := db.First(&kategori, newKuis.Kategori_id).Error; err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Invalid Kategori ID",
		})
	}

	// Lanjutkan dengan operasi database jika kategori valid
	result, err := database.CreateKuis(newKuis.Title, newKuis.Description, newKuis.Kategori_id, newKuis.Tingkatan_id, newKuis.Kelas_id)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Task added!",
	})
}

func UpdateKuis(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(500).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
	}

	newTask := new(models.Kuis)
	err = c.BodyParser(newTask)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.UpdateKuis(newTask.Title, newTask.Description, newTask.Kategori_id, newTask.Tingkatan_id, newTask.Kelas_id, id)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Task Updated!",
	})
	return nil
}

func DeleteKuis(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(500).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
	}

	err = database.DeleteKuis(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data":    nil,
		"success": true,
		"message": "Task Deleted Successfully",
	})
}
