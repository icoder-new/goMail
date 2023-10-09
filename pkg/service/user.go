package service

import (
	"fr33d0mz/goMail/logger"
	"fr33d0mz/goMail/models"
	"fr33d0mz/goMail/pkg/repository"
	"fr33d0mz/goMail/utils"
	"time"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: *repo,
	}
}

func (u *UserService) CreateUser(firstname, lastname, email string) (int, error) {
	var user models.User

	user.Firstname = firstname
	user.Lastname = lastname
	user.Email = email
	user.CreatedAt = time.Now()

	logger.Debug.Printf("{'user': %s %s %s %s}\n", user.Firstname, user.Lastname, user.Email, user.CreatedAt)
	return u.repo.CreateUser(user)

}

func (u *UserService) SendMessage(email, subject, message string) error {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	err = utils.SendMessage(email, user.Firstname, subject, message)
	if err != nil {
		return err
	}

	if err := u.repo.UpdateMessage(user.ID); err != nil {
		return err
	}

	logger.Debug.Println("Message was sent, yahoo")

	return nil
}

func (u *UserService) IsUserExistByEmail(email string) bool {
	return u.repo.IsUserExistByEmail(email)
}
