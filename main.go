package main

import (
	"github.com/Joko206/UAS_PWEB1/database"
	"github.com/Joko206/UAS_PWEB1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.GetDBConnection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		// Tentukan pengaturan CORS sesuai dengan kebutuhan
		AllowOrigins:     "http://10.0.2.2:8000",         // Misalnya hanya dari domain ini yang boleh mengakses
		AllowMethods:     "GET,POST,PUT,DELETE",          // Metode HTTP yang diizinkan
		AllowHeaders:     "Origin, Content-Type, Accept", // Header yang diizinkan
		AllowCredentials: true,                           // Mengizinkan pengiriman cookies
	}))

	routes.Setup(app)

	app.Listen("0.0.0.0:8000")

}
