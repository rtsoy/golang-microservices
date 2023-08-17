package scripts

import (
	"fmt"
	"log"

	"users/internal/domain"
	"users/internal/repository"

	"github.com/ddosify/go-faker/faker"
)

func Seed(repository *repository.Repository) {
	fake := faker.NewFaker()

	// Number of records
	n := 30

	// Seeding users into db
	for i := 0; i < n; i++ {
		user := &domain.User{
			FirstName:    fake.RandomPersonFirstName(),
			LastName:     fake.RandomPersonLastName(),
			Email:        fake.RandomEmail(),
			Password:     []byte("12345"),
			IsAmbassador: fake.RandomBoolean(),
		}
		user.HashPassword()

		if _, err := repository.CreateUser(user); err != nil {
			log.Fatalf("Failed to create a new record into db: %v", err)
		}
	}

	fmt.Printf("%d users were successfully added to db\n", n)
}
