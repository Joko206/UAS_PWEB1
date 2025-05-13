package main

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.GetDBConnection()

	app := fiber.New()

	routes.Setup(app)

	app.Listen("0.0.0.0:8000")

}
