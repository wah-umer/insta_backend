package main

import (
	"context"
	"fmt"
	"log"

	"github.com/umerwaheed/insta_backend/database"
)

func main() {
	database.ConnectDatabase()

	//check connection
	var name string
	err := database.DB.QueryRow(context.Background(), "SELECT current_database()").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current database:", name)

	defer database.DB.Close()
}
