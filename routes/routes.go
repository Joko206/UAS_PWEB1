package routes

import (
	"github.com/Joko206/UAS_PWEB1/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/user")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("haloo sekarang kamu ada didalam api user")
	})

	api.Get("/get-user", controllers.User)
	api.Post("/register", controllers.Register)
	api.Get("/login", controllers.Login)
	api.Get("/logout", controllers.Logout)

	kategori := app.Group("/kategori")

	kategori.Get("/", func(ctx *fiber.Ctx) error {
		_, err := controllers.Authenticate(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("haloo sekarang kamu ada didalam api kategori")
	})
	kategori.Get("/get-kategori", controllers.GetKategori)
	kategori.Post("/add-kategori", controllers.AddKategori)
	kategori.Patch("/update-kategori/:id", controllers.UpdateKategori)
	kategori.Delete("/delete-kategori/:id", controllers.DeleteKategori)

	tingkatan := app.Group("/tingkatan")

	tingkatan.Get("/", func(ctx *fiber.Ctx) error {
		_, err := controllers.Authenticate(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("haloo sekarang kamu ada didalam api kategori")
	})
	tingkatan.Get("/get-tingkatan", controllers.GetTingkatan)
	tingkatan.Post("/add-tingkatan", controllers.AddTingkatan)
	tingkatan.Patch("/update-tingkatan", controllers.UpdateTingkatan)
	tingkatan.Delete("/delete-tingkatan", controllers.DeleteTingkatan)

	kelas := app.Group("/kelas")

	kelas.Get("/", func(ctx *fiber.Ctx) error {
		_, err := controllers.Authenticate(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("haloo sekarang kamu ada didalam api kelass")
	})

	kelas.Get("/get-kelas", controllers.GetKelas)
	kelas.Post("/add-kelas", controllers.AddKelas)
	kelas.Patch("/update-kelas", controllers.UpdateKelas)
	kelas.Delete("/delete-kelas", controllers.DeleteKelas)

	Kuis := app.Group("/kuis")

	Kuis.Get("/", func(ctx *fiber.Ctx) error {
		_, err := controllers.Authenticate(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("haloo sekarang kamu ada didalam api Kuiss")
	})

	Kuis.Get("/get-kuis", controllers.GetKuis)
	Kuis.Post("/add-kuis", controllers.AddKuis)
	Kuis.Patch("/update-kuis", controllers.UpdateKuis)
	Kuis.Delete("/delete-kuis", controllers.DeleteKuis)

	Soal := app.Group("/soal")

	Soal.Get("/", func(ctx *fiber.Ctx) error {
		_, err := controllers.Authenticate(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("haloo sekarang kamu ada didalam api Soals")
	})

	Soal.Get("/get-soal", controllers.GetSoal)
	Soal.Post("/add-soal", controllers.AddSoal)
	Soal.Patch("/update-soal", controllers.UpdateSoal)
	Soal.Delete("/delete-soal", controllers.DeleteSoal)

}
