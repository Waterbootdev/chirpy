package database

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDatabaseQueries() *Queries {

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}

	return New(db)
}
