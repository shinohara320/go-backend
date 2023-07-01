package middleware

import "github.com/gofiber/fiber/v2"

func CorsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "OPTIONS" {
			c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Set("Access-Control-Allow-Credentials", "true")
			return c.SendStatus(fiber.StatusOK)
		}

		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Set("Access-Control-Allow-Credentials", "true")

		return c.Next()
	}
}
