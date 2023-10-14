package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PedroXimenes/4invest/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
	}
	rowsAffected, err := models.Delete(id)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("Could not find user: %d", id)
		return c.Status(http.StatusNotFound).SendString(errMsg)
	}
	msg := fmt.Sprintf("Foram excluídas %d linhas da tabela users.", rowsAffected)
	return c.Status(http.StatusOK).SendString(msg)
}
