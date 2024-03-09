package routes

import (
	"url-shortener/shortener"

	"github.com/gofiber/fiber/v2"
)

func Shortener(c *fiber.Ctx, shortener shortener.URLShortener) error {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	shortURL, err := shortener.HandleShorten(req.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"short_url": shortURL,
	})
}

func Short(c *fiber.Ctx, shortener shortener.URLShortener) error {
	shortKey := c.Params("shortKey")
	originalURL, err := shortener.HandleRedirect(shortKey)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(originalURL)
}
