package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"isAmbassador"`
}

func (u *User) HashPassword() {
	hash, _ := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	u.Password = hash
}

func (u *User) ComparePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(pw)) == nil
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	IsAmbassador    bool   `json:"isAmbassador"`
}
