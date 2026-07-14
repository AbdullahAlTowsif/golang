package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	return "user=postgres password=towsifsql host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbConnection, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbConnection, nil
}