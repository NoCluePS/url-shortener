package main

import (
	"url-shortener/routes"
	"url-shortener/shortener"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	shortener := &shortener.URLShortener{
		Urls: make(map[string]string),
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to URL Shortener")
	})

	app.Post("/shorten", func(c *fiber.Ctx) error { 
		return routes.Shortener(c, *shortener)
	})

	app.Get("/short/:shortKey", func(c *fiber.Ctx) error {
		return routes.Short(c, *shortener)
	})

	app.Listen(":8080")
}
