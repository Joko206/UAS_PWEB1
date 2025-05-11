package routes

import (
	"github.com/Joko206/UAS_PWEB1/controllers"
	"github.com/gofiber/fiber/v2"
)

// Middleware untuk autentikasi
func AuthMiddleware(c *fiber.Ctx) error {
	_, err := controllers.Authenticate(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    nil,
			"success": false,
			"message": "Unauthorized",
		})
	}
	return c.Next()
}

func Setup(app *fiber.App) {
	// Root Route
	start := app.Group("/")
	start.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	// User Routes
	api := app.Group("/user")
	api.Get("/get-user", controllers.User)
	api.Post("/register", controllers.Register)
	api.Get("/login", controllers.Login)
	api.Get("/logout", controllers.Logout)

	// Kategori Routes
	kategori := app.Group("/kategori", AuthMiddleware)
	kategori.Get("/get-kategori", controllers.GetKategori)
	kategori.Post("/add-kategori", controllers.AddKategori)
	kategori.Patch("/update-kategori/:id", controllers.UpdateKategori)
	kategori.Delete("/delete-kategori/:id", controllers.DeleteKategori)

	// Tingkatan Routes
	tingkatan := app.Group("/tingkatan", AuthMiddleware)
	tingkatan.Get("/get-tingkatan", controllers.GetTingkatan)
	tingkatan.Post("/add-tingkatan", controllers.AddTingkatan)
	tingkatan.Patch("/update-tingkatan", controllers.UpdateTingkatan)
	tingkatan.Delete("/delete-tingkatan", controllers.DeleteTingkatan)

	// Kelas Routes
	kelas := app.Group("/kelas", AuthMiddleware)
	kelas.Get("/get-kelas", controllers.GetKelas)
	kelas.Post("/add-kelas", controllers.AddKelas)
	kelas.Patch("/update-kelas", controllers.UpdateKelas)
	kelas.Delete("/delete-kelas", controllers.DeleteKelas)
	kelas.Post("/join-kelas", controllers.JoinKelas)

	// Kuis Routes
	kuis := app.Group("/kuis", AuthMiddleware)
	kuis.Get("/get-kuis", controllers.GetKuis)
	kuis.Post("/add-kuis", controllers.AddKuis)
	kuis.Patch("/update-kuis", controllers.UpdateKuis)
	kuis.Delete("/delete-kuis", controllers.DeleteKuis)

	// Soal Routes
	soal := app.Group("/soal", AuthMiddleware)
	soal.Get("/get-soal", controllers.GetSoal)
	soal.Post("/add-soal", controllers.AddSoal)
	soal.Patch("/update-soal", controllers.UpdateSoal)
	soal.Delete("/delete-soal", controllers.DeleteSoal)

	// Pendidikan Routes
	pendidikan := app.Group("/pendidikan", AuthMiddleware)
	pendidikan.Get("/get-pendidikan", controllers.GetPendidikan)
	pendidikan.Post("/add-pendidikan", controllers.AddPendidikan)
	pendidikan.Patch("/update-pendidikan", controllers.UpdatePendidikan)
	pendidikan.Delete("/delete-pendidikan", controllers.DeletePendidikan)

	// Hasil Kuis Routes
	result := app.Group("/hasil-kuis", AuthMiddleware)
	result.Get("/:user_id/:kuis_id", controllers.GetHasilKuis)
	result.Post("/submit-jawaban", controllers.SubmitJawaban)
}
