package handlers

import (
	"fmt"
	"net/http"

	"github.com/PedroXimenes/4invest/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Authorize(c *fiber.Ctx) error {
	user := &models.User{}
	err := c.BodyParser(user)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not decode request body")
	}
	key, err := user.ValidateAuthReq()
	if err != nil {
		log.WithField("key", key).Error("Missing key")
		errMsg := fmt.Sprintf("Missing key: %s", key)
		return c.Status(http.StatusBadRequest).SendString(errMsg)
	}

	if err := models.Authorize(user); err != nil {
		if err.Error() == "Incorrect email or password" {
			log.Error(err)
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		} else if err.Error() == "sql: no rows in result set" {
			log.Error(err)
			return c.Status(http.StatusUnauthorized).SendString("Incorrect email or password")
		} else {
			log.Error(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
	}
	return c.Status(http.StatusOK).SendString("ok")
}
