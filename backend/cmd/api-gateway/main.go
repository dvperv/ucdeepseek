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

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Заглушки для маршрутов
	// В дальнейшем здесь будет проксирование к микросервисам
	app.Get("/api/v1/users/*", proxyToUserService)
	app.Post("/api/v1/users/*", proxyToUserService)
	app.Get("/api/v1/communications/*", proxyToCommunicationService)
	app.Post("/api/v1/communications/*", proxyToCommunicationService)

	// Запуск сервера
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	app.Shutdown()
}

func proxyToUserService(c *fiber.Ctx) error {
	// Заглушка: возвращаем JSON с сообщением
	return c.JSON(fiber.Map{"service": "user", "path": c.Path()})
}

func proxyToCommunicationService(c *fiber.Ctx) error {
	// Заглушка: возвращаем JSON с сообщением
	return c.JSON(fiber.Map{"service": "communication", "path": c.Path()})
}
