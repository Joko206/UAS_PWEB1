package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize database connection
	_, err := database.GetDBConnection()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Create Fiber app with configuration
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		},
		ServerHeader: "BrainQuiz API",
		AppName:      "BrainQuiz v1.0",
	})

	// Add security middleware
	app.Use(helmet.New())
	app.Use(recover.New())

	// Add logging middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Add rate limiting
	app.Use(limiter.New(limiter.Config{
		Max: 100, // 100 requests per minute
	}))

	// Configure CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: getEnv("CORS_ORIGINS", "http://localhost:5173,https://brainquiz-psi.vercel.app"),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		}, ","),
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "BrainQuiz API is running",
			"version": "1.0.0",
		})
	})

	// Setup routes
	routes.Setup(app)

	// Get port from environment or use default
	port := getEnv("PORT", "8000")

	// Log environment info for debugging
	log.Printf("🌍 Environment: %s", getEnv("RAILWAY_ENVIRONMENT", "development"))
	log.Printf("🔗 Database URL configured: %t", os.Getenv("DATABASE_URL") != "")

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("🔄 Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// Start server
	log.Printf("🚀 Server starting on 0.0.0.0:%s", port)
	log.Printf("🏥 Health check available at: http://0.0.0.0:%s/health", port)

	if err := app.Listen("0.0.0.0:" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
