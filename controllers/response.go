package controllers

import "github.com/gofiber/fiber/v2"

// Helper function untuk membuat response JSON secara konsisten
func sendResponse(c *fiber.Ctx, status int, success bool, message string, data interface{}) error {
	return c.Status(status).JSON(&fiber.Map{
		"data":    data,
		"success": success,
		"message": message,
	})
}

// Helper function untuk menangani error
func handleError(c *fiber.Ctx, err error, message string) error {
	return sendResponse(c, fiber.StatusInternalServerError, false, message, err.Error())
}
