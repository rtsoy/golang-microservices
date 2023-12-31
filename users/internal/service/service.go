package service

import (
	"users/internal/domain"
	"users/internal/repository"
)

type Service struct {
	UserService
}

func NewService(rpstry *repository.Repository) *Service {
	return &Service{
		UserService: newUserService(rpstry),
	}
}

type UserService interface {
	GenerateToken(input domain.LoginInput) (string, error)
	CreateUser(input domain.RegisterInput) (*domain.User, error)
}
