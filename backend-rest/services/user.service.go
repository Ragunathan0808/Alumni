package service

import "github.com/kabi175/alumini-app-backend/domain/models"

type DefaultUserService struct {
	ur models.UserRepo
}

func NewDefaultUserService(ur models.UserRepo) models.UserService {
	return &DefaultUserService{
		ur: ur,
	}
}

func (as *DefaultUserService) GetByID(userID uint) (*models.User, error) {
	user, err := as.ur.FindByID(userID)
	return user, err
}

func (as *DefaultUserService) GetUser(user *models.User) ([]*models.User, error) {
	users, err := as.ur.Find(user)
	return users, err
}

func (as *DefaultUserService) GetByEmail(email string) (*models.User, error) {
	user, err := as.ur.FindByMail(email)
	return user, err
}

func (as *DefaultUserService) GetByDesignation(designation string) ([]*models.User, error) {
	users, err := as.ur.FindByDesignation(designation)
	return users, err
}

func (as *DefaultUserService) GetByBatch(batch uint) ([]*models.User, error) {
	users, err := as.ur.FindByBatch(batch)
	return users, err
}
