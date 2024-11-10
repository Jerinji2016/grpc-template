package repositories

import (
	"errors"

	"github.com/Jerinji2016/grpc-template/src/internal/db"
	"github.com/Jerinji2016/grpc-template/src/internal/models"
)

type UserRepository struct{}

func (repo *UserRepository) CreateUser(user *models.User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return errors.New("username already taken")
	}
	return nil
}

func (repo *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}