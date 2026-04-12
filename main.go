package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	app := fiber.New()
	var message []string
	app.Get("/start", func(c *fiber.Ctx) error {

		return c.SendString("Добро пожаловать в моё приложение FastAPI!")
	})

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendFile("index.html")
	})

	app.Get("/calculate", func(c *fiber.Ctx) error {
		num1 := c.Query("num1", "")
		num2 := c.Query("num2", "")
		num3, _ := strconv.Atoi(num1)
		num4, _ := strconv.Atoi(num2)
		num3 += num4
		num5 := strconv.Itoa(num3)
		otw := Otw{Result: num5}
		return c.JSON(otw)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		ot := User{Id: 0, Name: "qwerty"}
		return c.JSON(ot)
	})

	app.Post("/feedback", func(c *fiber.Ctx) error {
		ot := Feedback{}
		var response Response
		c.BodyParser(&ot)
		err := validate.Struct(ot)
		if err != nil {
			fmt.Println("Validation failed:", err)
			response.Message = "Использование недопустимых слов"
			return c.JSON(response)
		}
		message = append(message, ot.Message)

		response.Message = "Feedback received. Thank you, " + ot.Name + " ."
		return c.JSON(response)
	})

	app.Listen(":3000")
}
