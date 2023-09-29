package main

import (
	"fmt"
	"log"

	"github.com/PedroXimenes/4invest/configs"
	"github.com/PedroXimenes/4invest/internal/pkg/api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Could not load configurations from the config file.")
	}

	app := fiber.New()

	router.Router(app)

	port := fmt.Sprintf(":%s", configs.GetServerPort())
	err = app.Listen(port)
	if err != nil {
		log.Fatalf("Could not start server at port: %s\nError: %s", port, err.Error())
	}

}
