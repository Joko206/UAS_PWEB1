package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetTingkatan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetTingkatan()
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

func AddTingkatan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	newKategori := new(models.Kategori_Soal)
	err = c.BodyParser(newKategori)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.CreateTingkatan(newKategori.Name, newKategori.Description)
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
		"message": "Task added!",
	})
	return nil
}

func UpdateTingkatan(c *fiber.Ctx) error {
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

	newTask := new(models.Tingkatan)
	err = c.BodyParser(newTask)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.UpdateTingkatan(newTask.Name, newTask.Description, id)
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

func DeleteTingkatan(c *fiber.Ctx) error {
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

	err = database.DeleteTingkatan(id)
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
