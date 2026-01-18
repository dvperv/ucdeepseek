package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// WebSocket для чата
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/chat", websocket.New(handleWebSocket))

	// Маршруты для почты (заглушки)
	app.Get("/api/v1/emails", getEmails)
	app.Post("/api/v1/emails", sendEmail)

	// Запуск сервера на порту 3002
	go func() {
		if err := app.Listen(":3002"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	app.Shutdown()
}

func handleWebSocket(c *websocket.Conn) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func getEmails(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"emails": []string{"email1", "email2"}})
}

func sendEmail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Email sent"})
}
