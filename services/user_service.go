package services

import (
	"user-service/models"
	"user-service/repositories"
)

type UserService interface {
	Register(user *models.User) error
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
