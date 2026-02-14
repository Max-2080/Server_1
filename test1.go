package test1

import (
    "github.com/gofiber/fiber/v2"
    "fmt"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        
        return c.JSON({"message": "Добро пожаловать в моё
приложение FastAPI!"})
    })

    app.Listen(":3000")
}