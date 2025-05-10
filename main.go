package main

import (
	"belajar-via-dev.to/database"
	"belajar-via-dev.to/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DBconn()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")

}
