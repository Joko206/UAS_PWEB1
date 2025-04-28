package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Joko206/go_pisah/models"
)

func ProtectedHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User tidak ditemukan"})
		}

		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Halo %s!", user.Username),
			"user":    user,
		})
	}
}
