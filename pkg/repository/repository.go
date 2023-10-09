package repository

import (
	"fr33d0mz/goMail/models"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(user models.User) (int, error)
	UpdateMessage(id int) error
	IsUserExistByEmail(email string) bool
	GetUserByEmail(email string) (models.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
