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
	api.Post("/login", controllers.Login)
	api.Get("/logout", controllers.Logout)

	kategori := app.Group("/kategori", AuthMiddleware)
	kategori.Get("/get-kategori", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetKategori)
	kategori.Post("/add-kategori", controllers.RoleMiddleware([]string{"admin"}), controllers.AddKategori)
	kategori.Patch("/update-kategori/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.UpdateKategori)
	kategori.Delete("/delete-kategori/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.DeleteKategori)

	tingkatan := app.Group("/tingkatan", AuthMiddleware)
	tingkatan.Get("/get-tingkatan", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetTingkatan)
	tingkatan.Post("/add-tingkatan", controllers.RoleMiddleware([]string{"admin"}), controllers.AddTingkatan)
	tingkatan.Patch("/update-tingkatan/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.UpdateTingkatan)
	tingkatan.Delete("/delete-tingkatan/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.DeleteTingkatan)

	// Kelas Routes (Admin, Teacher, Student)
	kelas := app.Group("/kelas", AuthMiddleware)
	kelas.Get("/get-kelas", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetKelas)
	kelas.Post("/add-kelas", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.AddKelas)
	kelas.Patch("/update-kelas/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.UpdateKelas)
	kelas.Delete("/delete-kelas/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.DeleteKelas)
	kelas.Post("/join-kelas", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.JoinKelas)
	kelas.Get("/get-kelas-by-user", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetKelasByUserID)

	// Kuis Routes (Admin, Teacher)
	kuis := app.Group("/kuis", AuthMiddleware)
	kuis.Get("/get-kuis", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetKuis)
	kuis.Post("/add-kuis", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.AddKuis)
	kuis.Patch("/update-kuis/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.UpdateKuis)
	kuis.Delete("/delete-kuis/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.DeleteKuis)
	kuis.Get("/filter-kuis", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.FilterKuis)

	// Soal Routes (Admin, Teacher)
	soal := app.Group("/soal", AuthMiddleware)
	soal.Get("/get-soal", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetSoal)
	soal.Get("/get-soal/:kuis_id", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetSoalByKuisID)
	soal.Post("/add-soal", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.AddSoal)
	soal.Patch("/update-soal/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.UpdateSoal)
	soal.Delete("/delete-soal/:id", controllers.RoleMiddleware([]string{"admin", "teacher"}), controllers.DeleteSoal)

	// Pendidikan Routes (Only Admin)
	pendidikan := app.Group("/pendidikan", AuthMiddleware)
	pendidikan.Get("/get-pendidikan", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetPendidikan)
	pendidikan.Post("/add-pendidikan", controllers.RoleMiddleware([]string{"admin"}), controllers.AddPendidikan)
	pendidikan.Patch("/update-pendidikan/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.UpdatePendidikan)
	pendidikan.Delete("/delete-pendidikan/:id", controllers.RoleMiddleware([]string{"admin"}), controllers.DeletePendidikan)

	// Hasil Kuis Routes (Admin, Teacher, Student)
	result := app.Group("/hasil-kuis", AuthMiddleware)
	result.Get("/:user_id/:kuis_id", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.GetHasilKuis)
	result.Post("/submit-jawaban", controllers.RoleMiddleware([]string{"admin", "teacher, student"}), controllers.SubmitJawaban)
}
