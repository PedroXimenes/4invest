package models

import "github.com/PedroXimenes/4invest/internal/pkg/db"

func Get(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, email, name, age, nationality, created_at, updated_at FROM users WHERE id=$1`, id)

	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Age, &user.Nationality, &user.CreatedAt, &user.UpdatedAt)
	return
}
