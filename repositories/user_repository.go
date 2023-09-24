package repositories

import (
	"user-service/models"
	"user-service/pkg/db"
)

type UserRepository interface {
	Create(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserWithProfileById(userId uint) (*models.User, error)
}

type PGUserRepository struct{} // PostgreSQL
type MySQLRepository struct{}  // MySQL

func (repo *PGUserRepository) Create(user *models.User) error {
	return db.DB.Create(user).Error
}

func (repo *MySQLRepository) Create(user *models.User) error {
	return db.DB.Create(user).Error
}

func (repo *PGUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PGUserRepository) GetUserWithProfileById(userId uint) (*models.User, error) {
	var user models.User
	if err := db.DB.Preload("Profile").First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
