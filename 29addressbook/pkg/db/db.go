package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB
// var Contacts []model.Contact

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Fail to load .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db
}
