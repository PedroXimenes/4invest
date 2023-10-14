package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PedroXimenes/4invest/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Get(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
	}

	user, err := models.Get(id)
	if err != nil {
		log.Error(err)
		if err.Error() == "sql: no rows in result set" {
			errMsg := fmt.Sprintf("No results for user: %d", id)
			return c.Status(http.StatusNotFound).SendString(errMsg)
		}
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(http.StatusOK).JSON(user)
}
