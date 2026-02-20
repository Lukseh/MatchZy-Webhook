package util

import "github.com/gofiber/fiber/v3"

func AuthCheck(expected string) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if authHeader != expected {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
