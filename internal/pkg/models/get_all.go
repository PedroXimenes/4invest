package models

import "github.com/PedroXimenes/4invest/internal/pkg/db"

func GetAll() (users []User, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, email, name, age, nationality, created_at, updated_at FROM users`)
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Age, &user.Nationality, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			continue
		}

		users = append(users, user)

	}
	return
}
