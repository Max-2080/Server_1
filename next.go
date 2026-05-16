package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		login := Login{}
		c.BodyParser(&login)
		if login.Username == "user" {
			cookie := new(fiber.Cookie)
			cookie.Name = "session_token"                   // Имя cookie
			cookie.Value = "dark_mode"                      // Значение
			cookie.Expires = time.Now().Add(24 * time.Hour) // Время истечения
			cookie.HTTPOnly = true                          // Защита от XSS (JavaScript не сможет прочитать)
			cookie.Secure = false                           // Только по HTTPS (для production)
			cookie.SameSite = "Lax"                         // Политика для跨域请求
			cookie.Path = "/"                               // Доступно для всего сайта

			// 2. Отправляем cookie клиенту
			c.Cookie(cookie)
		}
		return c.SendStatus(200)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		sessionToken := c.Cookies("session_token")

		if sessionToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing cookie",
			})
		}
		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
