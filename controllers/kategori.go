package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
)

func GetKategori(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	result, err := database.GetKategori()
	if err != nil {
		return handleError(c, err, "Failed to fetch categories")
	}

	return sendResponse(c, fiber.StatusOK, true, "All categories retrieved successfully", result)
}

func AddKategori(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	newKategori := new(models.Kategori_Soal)
	if err := c.BodyParser(newKategori); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.CreateKategori(newKategori.Name, newKategori.Description)
	if err != nil {
		return handleError(c, err, "Failed to add category")
	}

	return sendResponse(c, fiber.StatusOK, true, "Category added successfully", result)
}

func UpdateKategori(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	newTask := new(models.Kategori_Soal)
	if err := c.BodyParser(newTask); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	result, err := database.UpdateKategori(newTask.Name, newTask.Description, id)
	if err != nil {
		return handleError(c, err, "Failed to update category")
	}

	return sendResponse(c, fiber.StatusOK, true, "Category updated successfully", result)
}

func DeleteKategori(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	_, err := Authenticate(c)
	if err != nil {
		return err
	}

	id := c.Params("id")
	if id == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "ID cannot be empty", nil)
	}

	err = database.DeleteKategori(id)
	if err != nil {
		return handleError(c, err, "Failed to delete category")
	}

	return sendResponse(c, fiber.StatusOK, true, "Category deleted successfully", nil)
}
