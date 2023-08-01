package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shinohara320/travel-agent/util"
)

func IsAuthenticate(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader != "" {
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		fmt.Println("Token from Authorization Header:", token)
		if _, err := util.ParseJwt(token); err == nil {
			return c.Next()
		}
	}

	return c.Next()
}
