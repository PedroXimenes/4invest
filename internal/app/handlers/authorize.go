package handlers

import (
	"fmt"
	"net/http"

	"github.com/PedroXimenes/4invest/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Authorize(c *fiber.Ctx) error {
	user := &models.User{}
	err := c.BodyParser(user)
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not decode request body")
	}
	key, err := user.ValidateInput()
	if err != nil {
		errMsg := fmt.Sprintf("Missing key: %s", key)
		return c.Status(http.StatusBadRequest).SendString(errMsg)
	}

	if err := models.Authorize(user); err != nil {
		if err.Error() == "Incorrect password" {
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		} else if err.Error() == "sql: no rows in result set" {
			return c.Status(http.StatusNotFound).SendString("Email not found")
		} else {
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
	}
	return c.Status(http.StatusOK).SendString("ok")
}
