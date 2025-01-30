package repository

import (
	"app"
	"database/sql"
	"fmt"
)

type MovePostgres struct {
	db *sql.DB
}

func NewMovePostgres(db *sql.DB) *MovePostgres {
	return &MovePostgres{db: db}
}

func (r *MovePostgres) CreateList(userTable app.UserTable) (int, error) {
	stmt, err := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id", USERSTABLE))
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	var id int
	row := stmt.QueryRow(userTable.UserId, userTable.Title, userTable.Description)
	if err = row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to scan id: %v", err)
	}

	return id, nil
}

func (r *MovePostgres) GetListsUser(userId int) ([]app.UserTable, error) {
	var result []app.UserTable
	stmt, err := r.db.Prepare(fmt.Sprintf("SELECT id, title, description FROM %s WHERE user_id = $1", USERSTABLE))
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}
	for rows.Next() {
		var userTable app.UserTable
		if err = rows.Scan(&userTable.ListId, &userTable.Title, &userTable.Description); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		result = append(result, userTable)
	}
	return result, nil
}
