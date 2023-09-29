package models

import "github.com/PedroXimenes/4invest/internal/pkg/db"

func GetAll() (users []User, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			continue
		}

		users = append(users, user)

	}
	return
}
