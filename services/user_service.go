package services

import (
	"user-service/models"
	"user-service/repositories"
)

type UserService interface {
	Register(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserWithProfileById(userId uint) (*models.User, error)
}

type DefaultUserService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &DefaultUserService{repo: r}
}

func (svc *DefaultUserService) Register(user *models.User) error {
	return svc.repo.Create(user)
}

func (svc *DefaultUserService) GetUserByEmail(email string) (*models.User, error) {
	return svc.repo.GetUserByEmail(email)
}

func (svc *DefaultUserService) GetUserWithProfileById(userId uint) (*models.User, error) {
	return svc.repo.GetUserWithProfileById(userId)
}
