package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
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
func verifyAndExtractUserID(sessionToken string) (string, int, bool) {
	parts := strings.Split(sessionToken, ".")
	if len(parts) != 3 {
		return "", 0, false
	}

	userID := parts[0]
	times, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	receivedSignature := parts[2]

	// Перевычисляем подпись
	expectedSignature := generateSignature(userID)

	// Сравниваем (constant-time comparison для безопасности)
	return userID, times, hmac.Equal([]byte(receivedSignature), []byte(expectedSignature))
}

func main() {
	app := fiber.New()

	var Users = []Login{}

	app.Post("/login", func(c *fiber.Ctx) error {
		login := Login{}
		c.BodyParser(&login)
		c.Cookie(&fiber.Cookie{
			Name:     "session_token",
			Value:    login.Username + "." + strconv.Itoa(int(time.Now().Unix())) + "." + generateSignature(login.Username),
			HTTPOnly: true,
			//Expires:  expires,
			MaxAge: 300,
			Path:   "/",
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

		userid, date, err := verifyAndExtractUserID(sessionToken)
		if err == false {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}
		if date+180 < int(time.Now().Unix()) {
			parts := strings.Split(c.Cookies("session_token"), ".")
			fmt.Println(parts)
			c.Cookie(&fiber.Cookie{
				Name:     "session_token",
				Value:    parts[0] + "." + strconv.Itoa(int(time.Now().Unix())) + "." + generateSignature(parts[0]), // новое значение
				MaxAge:   300,
				HTTPOnly: true,
				Path:     "/",
			})
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

	app.Get("/headers", func(c *fiber.Ctx) error {
		user_agent := c.Get("User-agent")
		accept_language := c.Get("Accept-Language")
		commonheaders := CommonHeaders{}
		commonheaders.Accept_Language = accept_language
		commonheaders.User_Agent = user_agent
		if user_agent == "" || accept_language == "" {
			return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"message": "NEED HEADERS!"})
		}
		return c.JSON(commonheaders)
	})

	app.Get("/info", func(c *fiber.Ctx) error {
		user_agent := c.Get("User-agent")
		accept_language := c.Get("Accept-Language")
		commonheaders := CommonHeaders{}
		commonheaders.Accept_Language = accept_language
		commonheaders.User_Agent = user_agent
		if user_agent == "" || accept_language == "" {
			return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"message": "NEED HEADERS!"})
		}
		c.Set("X-Server-Time", time.Now().Format(time.RFC3339))
		return c.JSON(fiber.Map{"CommonHeaders": commonheaders,
			"message": " Добро пожаловать! Ваши заголовки успешно обработаны."})
	})

	app.Listen(":3000")
}
