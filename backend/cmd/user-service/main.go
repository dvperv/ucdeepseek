package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Маршруты для пользователей
	app.Get("/api/v1/users", getUsers)
	app.Get("/api/v1/users/:id", getUser)
	app.Post("/api/v1/users", createUser)
	app.Put("/api/v1/users/:id", updateUser)
	app.Delete("/api/v1/users/:id", deleteUser)

	// Запуск сервера на порту 3001
	go func() {
		if err := app.Listen(":3001"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	app.Shutdown()
}

// Обработчики маршрутов (заглушки)
func getUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"users": []string{"user1", "user2"}})
}

func getUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": c.Params("id"), "name": "John Doe"})
}

func createUser(c *fiber.Ctx) error {
	return c.Status(201).JSON(fiber.Map{"message": "User created"})
}

func updateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": c.Params("id"), "updated": true})
}

func deleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": c.Params("id"), "deleted": true})
}
