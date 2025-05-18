package controllers

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

// Secret key for JWT
const SecretKey = "secret"

// Helper function to authenticate using JWT
func Authenticate(c *fiber.Ctx) (*models.Users, error) {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "No JWT token found")
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.Users
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	if user.ID == 0 {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	return &user, nil
}

func RoleMiddleware(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := Authenticate(c)
		if err != nil {
			return err
		}

		// Check if the user role is allowed
		roleAllowed := false
		for _, role := range allowedRoles {
			if user.Role == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"data":    nil,
				"success": false,
				"message": "You don't have permission to access this resource",
			})
		}

		return c.Next()
	}
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// Parse request body
	if err := c.BodyParser(&data); err != nil {
		log.Printf("Error parsing body: %v", err)
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	// Validasi jika name, email, atau password kosong
	if data["name"] == "" || data["email"] == "" || data["password"] == "" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Name, email, and password are required", nil)
	}

	// Default role adalah "student" jika tidak ada
	role := data["role"]
	if role == "" {
		role = "student" // Role default
	}

	// Validasi role yang diizinkan
	if role != "admin" && role != "teacher" && role != "student" {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid role. Allowed roles: admin, teacher, student", nil)
	}

	// Hash password sebelum disimpan
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return sendResponse(c, fiber.StatusInternalServerError, false, "Error hashing password", nil)
	}

	// Buat user baru
	user := models.Users{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
		Role:     role,
	}

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Failed to register user: %v", err)
		return handleError(c, err, "Failed to register user")
	}

	// Kirim response sukses
	return sendResponse(c, fiber.StatusOK, true, "User registered successfully", user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// Parse request body
	if err := c.BodyParser(&data); err != nil {
		return sendResponse(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	var user models.Users
	// Find user by email
	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
		return sendResponse(c, fiber.StatusNotFound, false, "User not found", nil)
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return sendResponse(c, fiber.StatusUnauthorized, false, "Invalid password", nil)
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return handleError(c, err, "Failed to generate token")
	}

	// Set JWT cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return sendResponse(c, fiber.StatusOK, true, "Login successful", fiber.Map{
		"token": token,
		"role":  user.Role,
	})
}

func User(c *fiber.Ctx) error {
	// Authenticate the user using the JWT token
	user, err := Authenticate(c)
	if err != nil {
		return err
	}

	// Return user details
	return sendResponse(c, fiber.StatusOK, true, "User retrieved successfully", user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	// Return success message
	return sendResponse(c, fiber.StatusOK, true, "Logout successful", nil)
}
