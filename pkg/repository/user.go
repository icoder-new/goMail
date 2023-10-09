package repository

import (
	"fr33d0mz/goMail/logger"
	"fr33d0mz/goMail/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user models.User) (int, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		logger.Debug.Printf("[REPO]|CreateUser: %s\n", err)
		return -1, err
	}

	logger.Debug.Println("[REPO]: user created")

	return user.ID, nil
}

func (u *UserRepository) UpdateMessage(id int) error {
	var user models.User

	err := u.db.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Debug.Printf("[REPO]|UpdateMessage: %s\n", err)
		return err
	}

	user.MessageSentCount += 1
	user.MessageSentDate = time.Now()

	err = u.db.Updates(&user).Error
	if err != nil {
		logger.Debug.Printf("[REPO]|UpdateMessage(2): %s\n", err)
		return err
	}

	return nil
}

func (u *UserRepository) IsUserExistByEmail(email string) bool {
	var user models.User
	err := u.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Debug.Printf("[repo|err]IsUserExistByEmail: %s\n", err)
		return false
	}

	return true
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Debug.Printf("[REPO]|GetUserByEmail: %s\n", err)
		return models.User{}, err
	}

	return user, nil
}
