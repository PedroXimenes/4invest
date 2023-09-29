package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        int64
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string
	UpdatedAt string
}

func AddTimestamp() (string, error) {
	localZone, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return "", err
	}
	now := time.Now().In(localZone).Format("2006-01-02 15:04:05-07:00")

	return now, nil
}

func (u *User) ValidateInput() (key string, err error) {
	if u.Email == "" {
		err = fmt.Errorf("Missing key: email")
		key = "email"
	} else if u.Password == "" {
		err = fmt.Errorf("Missing key: password")
		key = "password"
	}
	return
}
