package middleware

import "github.com/gofiber/fiber/v2"

// CorsMiddleware adalah middleware CORS untuk mengatur akses CORS pada server
func CorsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "OPTIONS" {
			// Tangani permintaan Preflight
			c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Set("Access-Control-Allow-Credentials", "true")
			return c.SendStatus(fiber.StatusOK)
		}

		// Lanjutkan ke middleware berikutnya
		return c.Next()
	}
}
