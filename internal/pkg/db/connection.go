package db

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/PedroXimenes/4invest/configs"

// 	_ "github.com/lib/pq"
// )

// func OpenConnection() (*sql.DB, error) {
// 	conf := configs.GetDB()

// 	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
// 		conf.Host,
// 		conf.Port,
// 		conf.User,
// 		conf.Pass,
// 		conf.Database)

// 	conn, err := sql.Open("postgres", sc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = conn.Ping()

// 	return conn, err
// }

import (
	"database/sql"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func OpenConnection() (*sql.DB, error) {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
		return nil, err
	}
	defer cleanup()

	db, err := sql.Open(
		"cloudsql-postgres",
		"host=sincere-almanac-401411:us-central1:four-invest user=postgres password=m-,b(<(vIgsuXCV< dbname=four-invest sslmode=disable",
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
