package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	errLoad := godotenv.Load(".env")

	if errLoad != nil {
		log.Fatal("Error loading .env file")
	}

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("DB"),
	)

	var err error
	Client, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
