package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Joko206/go_pisah/handlers"
	"github.com/Joko206/go_pisah/middleware"
	"github.com/Joko206/go_pisah/models"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=joko port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database")
	}
	db.AutoMigrate(&models.User{})

	app := fiber.New()
	secret := []byte(os.Getenv("JWT_SECRET"))

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool { return true },
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Post("/register", handlers.RegisterHandler(db))
	app.Post("/login", handlers.LoginHandler(db, secret))
	app.Get("/protected", middleware.AuthMiddleware(secret), handlers.ProtectedHandler(db))

	app.Listen(":3000")
}
