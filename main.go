package main

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func main() {

	database.GetDBConnection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:5173",
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
