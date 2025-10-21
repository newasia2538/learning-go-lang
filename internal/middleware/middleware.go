package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("\nURL : %s, Method : %s, timestamp : %s", c.OriginalURL(), c.Method(), start)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
