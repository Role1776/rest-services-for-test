package service

import (
	"app"
)

//go:generate mockgen -source=service.go -destination=mocks/mock_service.go -package=mocks

type Authorization interface {
	CreateUser(user app.User) (app.User, error)
	GeneratePasswordHash(password string) string
	ParseToken(token string) (int, error)
	GenerateJWT(id int) (string, error)
}

type Move interface {
	CreateList(userTable app.UserTable, id int) (int, error)
	GetListsUser(id int) ([]app.UserTable, error)
}

type Service struct {
	Authorization
	Move
}

func NewService() *Service {
	return &Service{
		Authorization: &NewAuthorization{},
	}
}
