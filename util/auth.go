package util

import "github.com/gofiber/fiber/v3"

func AuthCheck(auth struct {
	Header string `json:"header"`
	Value  string `json:"value"`
}) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		authHeader := c.Get(auth.Header)

		if authHeader == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if authHeader != auth.Value {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
