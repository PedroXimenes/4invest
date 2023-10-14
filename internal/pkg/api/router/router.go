package router

import (
	"github.com/PedroXimenes/4invest/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://four-invest-front-p3xh7jp6wa-uc.a.run.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/users", handlers.GetAll)
	app.Get("/users/:id", handlers.Get)
	app.Post("/users", handlers.Create)
	app.Post("/users/auth", handlers.Authorize)
	app.Delete("/users/:id", handlers.Delete)
	//app.Put("/users/update/:id/username", handlers.Update)
}
