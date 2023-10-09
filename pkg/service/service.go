package service

import "fr33d0mz/goMail/pkg/repository"

type User interface {
	CreateUser(firstname, lastname, email string) (int, error)
	SendMessage(email, subject, message string) error
	IsUserExistByEmail(email string) bool
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
