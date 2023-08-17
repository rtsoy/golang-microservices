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

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if tx := r.db.Create(user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
