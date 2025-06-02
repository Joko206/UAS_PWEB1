package controllers

import (
	"fmt"

	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetKelas()
	if err != nil {
		return handleError(c, err, "Failed to retrieve classes")
	}

	return sendResponse(c, fiber.StatusOK, true, "All classes retrieved successfully", result)
}

func AddKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	newKategori := new(models.Kelas)
	if err := c.BodyParser(newKategori); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.CreateKelas(newKategori.Name, newKategori.Description)
	if err != nil {
		return handleError(c, err, "Failed to add class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class added successfully", result)
}

func UpdateKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	newTask := new(models.Kelas)
	if err := c.BodyParser(newTask); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdateKelas(newTask.Name, newTask.Description, id)
	if err != nil {
		return handleError(c, err, "Failed to update class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class updated successfully", result)
}

func DeleteKelas(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err = database.DeleteKelas(id)
	if err != nil {
		return handleError(c, err, "Failed to delete class")
	}

	return sendResponse(c, fiber.StatusOK, true, "Class deleted successfully", nil)
}

func GetKelasById(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		fmt.Printf("Authentication error: %v\n", err)
		return err
	}

	id := c.Params("id")
	if id == "" {
		fmt.Println("ID parameter is empty")
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	fmt.Printf("Fetching class with ID: %s\n", id)

	result, err := database.GetKelasById(id)
	if err != nil {
		fmt.Printf("Error fetching class: %v\n", err)
		return handleError(c, err, "Failed to retrieve class")
	}

	fmt.Printf("Successfully retrieved class: %+v\n", result)
	return sendResponse(c, fiber.StatusOK, true, "Class retrieved successfully", result)
}
