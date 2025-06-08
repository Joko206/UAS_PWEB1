package main

import (
	"log"
	"os"

	"strings"

	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Check if seed argument is provided
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		log.Println("Running database seeding...")

		// Initialize database connection
		_, err := database.InitDB()
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}

		// Run the seeding
		if err := database.SeedDatabase(); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}

		log.Println("Database seeding completed successfully!")
		return
	}

	database.GetDBConnection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, https://brainquiz-psi.vercel.app/, https://brainquizz1.vercel.app/",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen("0.0.0.0:8000")

}
