package repositories

import (
	"user-service/models"
	"user-service/pkg/db"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type PGUserRepository struct{} // PostgreSQL
type MySQLRepository struct{}  // MySQL

func (repo *PGUserRepository) Create(user *models.User) error {
	return db.DB.Create(user).Error
}

func (repo *MySQLRepository) Create(user *models.User) error {
	return db.DB.Create(user).Error
}

func (repo *PGUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
