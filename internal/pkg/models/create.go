package models

import (
	"crypto/sha256"
	"fmt"

	"github.com/PedroXimenes/4invest/internal/pkg/db"
)

func Insert(user *User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Printf("DB Connection error: %s\n", err)
		return
	}
	defer conn.Close()

	user.hashPassword()

	now, err := AddTimestamp()
	if err != nil {
		fmt.Printf("Timestamp error: %s\n", err)
		return
	}

	sql := `INSERT INTO users (email, password, name, age, nationality, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = conn.QueryRow(sql, user.Email, user.Password, user.Name, user.Age, user.Nationality, now, now).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (u *User) hashPassword() {
	hashed := sha256.Sum256([]byte(addSalt(u.Password)))
	u.Password = fmt.Sprintf("%x", hashed)
}

func addSalt(pass string) string {
	return fmt.Sprintf("investsalt%s", pass)
}
