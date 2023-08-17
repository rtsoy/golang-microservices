package main

import (
	"log"

	"users/pkg/mysql"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	_ = db

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}
