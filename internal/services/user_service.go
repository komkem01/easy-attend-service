package services

import (
	"easy-attend-service/internal/database"
	"easy-attend-service/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(user *models.User) error {
	return database.GetDB().Create(user).Error
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.GetDB().First(&user, id).Error
	return &user, err
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.GetDB().Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *UserService) UpdateUser(user *models.User) error {
	return database.GetDB().Save(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return database.GetDB().Delete(&models.User{}, id).Error
}
