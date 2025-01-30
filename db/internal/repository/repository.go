package repository

import (
	"app"
	"database/sql"
)

type Authorization interface {
	CreateUser(app.User) (int, error)
	GetUser(username, password string) (int, error)
}

type Move interface {
	CreateList(userTable app.UserTable) (int, error)
	GetListsUser(userId int) ([]app.UserTable, error)
}

type Repository struct {
	Authorization
	Move
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Authorization: NewAuthPostgres(db),
		Move:          NewMovePostgres(db),
	}
}
