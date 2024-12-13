package database

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// DBInit initializes and returns the database connection with retry logic
func DBInit() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "postgres"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "careem"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	var db *gorm.DB
	var err error

	// Retry logic
	for i := 0; i < 5; i++ {
		db, err = gorm.Open("postgres", dsn)
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database (attempt %d/5): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("could not connect to database after 5 attempts: %w", err)
	}

	return db, nil
}
