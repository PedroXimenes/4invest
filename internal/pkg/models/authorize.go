package models

import (
	"fmt"

	"github.com/PedroXimenes/4invest/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

func Authorize(user *User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Errorf("DB Connection error: %s\n", err)
		return 0, err
	}
	defer conn.Close()

	user.hashPassword()

	row := conn.QueryRow(`SELECT id, password  FROM users WHERE email=$1`, user.Email)
	type pass struct {
		password string
	}
	p := &pass{}
	var id int64
	err = row.Scan(&id, &p.password)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	if p.password != user.Password {
		err := fmt.Errorf("Incorrect email or password")
		return 0, err
	}

	return id, nil
}
