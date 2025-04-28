package handlers

import (
	"github.com/Joko206/go_pisah/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func RegisterHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.RegisterInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal hash password"})
		}

		user := models.User{
			Username: input.Username,
			Email:    input.Email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username atau email sudah terdaftar"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Registrasi berhasil"})
	}
}

func LoginHandler(db *gorm.DB, secret []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username tidak ditemukan"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Password salah"})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(secret)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat token"})
		}

		return c.JSON(fiber.Map{"token": tokenString})
	}
}
