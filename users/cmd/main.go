package main

import (
	"log"

	"users/internal/handler"
	"users/internal/repository"
	"users/internal/service"
	"users/pkg/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := mysql.New()
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	var (
		rpstry = repository.NewRepository(db)
		svc    = service.NewService(rpstry)
		hndlr  = handler.NewHandler(svc)

		app = fiber.New()
	)

	hndlr.InitRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":8000")
}
