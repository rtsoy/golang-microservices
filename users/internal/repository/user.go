package repository

import (
	"users/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if tx := r.db.Where("email = ?", email).First(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if tx := r.db.Create(user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
