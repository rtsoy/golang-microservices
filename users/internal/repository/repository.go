package repository

import (
	"users/internal/domain"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: newUserRepository(db),
	}
}

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
}
