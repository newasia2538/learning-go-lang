package restapi

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type User = struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var defaultUser = User{
	Email:    "user@example.com",
	Password: "Password123",
}

func login(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != defaultUser.Email || user.Password != defaultUser.Password {
		return fiber.ErrUnauthorized
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login success !",
		"token":   t,
	})
}
