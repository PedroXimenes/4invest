package router

import (
	"github.com/PedroXimenes/4invest/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/users", handlers.GetAll)
	app.Get("/users/:id", handlers.Get)
	app.Post("/users", handlers.Create)
	app.Post("/users/auth", handlers.Authorize)
	app.Delete("/users/:id", handlers.Delete)
	//app.Put("/users/update/:id/username", handlers.Update)
}
