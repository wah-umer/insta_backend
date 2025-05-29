package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_string := os.Getenv("DB_URL")

	pool, err := pgxpool.New(context.Background(), db_string)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	DB = pool
	log.Println("Connected to PostgreSQL with pgxpool!")
}
