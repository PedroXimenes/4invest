package router

import (
	"github.com/PedroXimenes/4invest/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/users", handlers.GetAll)
	app.Get("/users/:id", handlers.Get)
	app.Post("/users", handlers.Create)
	app.Post("/users/auth", handlers.Authorize)
	app.Delete("/users/:id", handlers.Delete)
	//app.Put("/users/update/:id/username", handlers.Update)
}
