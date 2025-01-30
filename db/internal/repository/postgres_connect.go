package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewConn() (*sql.DB, error) {
	connParam := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5555", "postgres", "1", "postgres", "disable")
	db, err := sql.Open("pgx", connParam)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Проверяем существование таблицы users
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "users").Scan(&exists)
	if err != nil {
		log.Printf("Error checking table existence: %v", err)
		return nil, fmt.Errorf("failed to check table existence: %v", err)
	}

	if !exists {
		log.Printf("Table 'users' does not exist, creating...")
		_, err = db.Exec(`
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) NOT NULL UNIQUE,
				password_hash VARCHAR(255) NOT NULL
			)
		`)
		if err != nil {
			log.Printf("Error creating table: %v", err)
			return nil, fmt.Errorf("failed to create table: %v", err)
		}
		log.Printf("Table 'users' created successfully")
	} else {
		log.Printf("Table 'users' already exists")
	}

	return db, nil
}
