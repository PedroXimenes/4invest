package handlers

import (
	"fmt"
	"net/http"

	"github.com/PedroXimenes/4invest/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	user := &models.User{}
	err := c.BodyParser(user)
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(http.StatusInternalServerError).SendString("Could not decode request body")
	}
	key, err := user.ValidateInput()
	if err != nil {
		errMsg := fmt.Sprintf("Missing key: %s", key)
		return c.Status(http.StatusBadRequest).SendString(errMsg)
	}
	id, err := models.Insert(user)
	if err != nil {
		fmt.Printf("%v\n", err)
		if err.Error() == `pq: duplicate key value violates unique constraint "unique_email"` {
			return c.Status(http.StatusConflict).SendString("This email is already in use")
		} else if err.Error() == `pq: duplicate key value violates unique constraint "unique_username"` {
			return c.Status(http.StatusConflict).SendString("This username is already in use")
		}
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	message := fmt.Sprintf("Created with ID: %d.", id)
	return c.Status(http.StatusCreated).SendString(message)
}
