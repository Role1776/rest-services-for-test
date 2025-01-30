package repository

import (
	"app"
	"database/sql"
	"fmt"
	"log"
)

const (
	USERS      = "users"
	USERSTABLE = "user_tables"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int

	log.Printf("Attempting to create user: %s", user.Username)

	stmt, err := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", USERS))
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}

	row := stmt.QueryRow(user.Username, user.Password)
	if err = row.Scan(&id); err != nil {
		log.Printf("Error scanning id: %v", err)
		return 0, fmt.Errorf("failed to scan id: %v", err)
	}

	log.Printf("User created successfully with id: %d", id)
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (int, error) {
	var id int

	log.Printf("Attempting to get user: %s", username)

	stmt, err := r.db.Prepare(fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", USERS))
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}

	row := stmt.QueryRow(username, password)

	if err = row.Scan(&id); err != nil {
		log.Printf("Error scanning id: %v", err)
		return 0, fmt.Errorf("user not found: %v", err)
	}

	log.Printf("User found with id: %d", id)
	return id, nil
}

//{
//"title":"купить бананы",
// "description":"бананы топ"
//}
