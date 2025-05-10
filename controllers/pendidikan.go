package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetPendidikan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetPendidikan()
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

func AddPendidikan(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	newKategori := new(models.Pendidikan)
	err = c.BodyParser(newKategori)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.CreatePendidikan(newKategori.Name, newKategori.Description)
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

func UpdatePendidikan(c *fiber.Ctx) error {
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

	newTask := new(models.Pendidikan)
	err = c.BodyParser(newTask)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.UpdatePendidikan(newTask.Name, newTask.Description, id)
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

func DeletePendidikan(c *fiber.Ctx) error {
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

	err = database.DeletePendidikan(id)
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
