package service

import (
	"errors"
	"net/mail"

	"users/internal/domain"
	"users/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "secret"

type userService struct {
	rpstry repository.UserRepository
}

func newUserService(rpstry repository.UserRepository) UserService {
	return &userService{
		rpstry: rpstry,
	}
}

func (s *userService) GenerateToken(input domain.LoginInput) (string, error) {
	user, err := s.rpstry.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	if !user.ComparePassword(input.Password) {
		return "", errors.New("Invalid Credentials")
	}

	claims := jwt.MapClaims{
		"email": input.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := SecretKey

	return token.SignedString([]byte(secret))
}

func (s *userService) CreateUser(input domain.RegisterInput) (*domain.User, error) {
	if input.Password != input.PasswordConfirm {
		return nil, errors.New("Passwords do not match!")
	}

	if _, err := mail.ParseAddress(input.Email); err != nil {
		return nil, errors.New("Invalid email")
	}

	if len(input.FirstName) < 0 {
		return nil, errors.New("First name cannot be empty")
	}

	if len(input.LastName) < 0 {
		return nil, errors.New("Last name cannot be empty")
	}

	user := &domain.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		Password:     []byte(input.Password),
		IsAmbassador: input.IsAmbassador,
	}
	user.HashPassword()

	return s.rpstry.CreateUser(user)
}
