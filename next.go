package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

var secretKey = []byte("my-super-secret-key-32-bytes-long!!!")

// Генерация подписи для user_id
func generateSignature(userID string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(userID))
	return hex.EncodeToString(h.Sum(nil))
}

// Проверка подписи и извлечение user_id
func verifyAndExtractUserID(sessionToken string) (string, bool) {
	parts := strings.Split(sessionToken, ":")
	if len(parts) != 2 {
		return "", false
	}

	userID := parts[0]
	receivedSignature := parts[1]

	// Перевычисляем подпись
	expectedSignature := generateSignature(userID)

	// Сравниваем (constant-time comparison для безопасности)
	return userID, hmac.Equal([]byte(receivedSignature), []byte(expectedSignature))
}

func main() {
	app := fiber.New()

	var Users = []Login{}

	app.Post("/login", func(c *fiber.Ctx) error {
		login := Login{}
		c.BodyParser(&login)

		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    "dark_mode",
			HTTPOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
		})
		Users = append(Users, login)
		return c.SendStatus(200)
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		sessionToken := c.Cookies("session_token")

		if sessionToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing cookie",
			})
		}

		userid, err := verifyAndExtractUserID(sessionToken)
		fmt.Println(userid)
		fmt.Println(sessionToken)
		if err != true {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}
		return c.JSON(userid)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		sessionToken := c.Cookies("session_token")
		fmt.Println(sessionToken)
		if sessionToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing cookie",
			})
		}
		return c.JSON(Users[0])
	})

	app.Listen(":3000")
}
