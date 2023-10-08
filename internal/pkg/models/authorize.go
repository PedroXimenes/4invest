package models

import (
	"fmt"

	"github.com/PedroXimenes/4invest/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

func Authorize(user *User) error {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Errorf("DB Connection error: %s\n", err)
		return err
	}
	defer conn.Close()

	user.hashPassword()

	row := conn.QueryRow(`SELECT password FROM users WHERE email=$1`, user.Email)
	type pass struct {
		password string
	}
	p := &pass{}
	err = row.Scan(&p.password)
	if err != nil {
		log.Error(err)
		return err
	}

	if p.password != user.Password {
		err := fmt.Errorf("Incorrect password")
		return err
	}

	return nil
}
