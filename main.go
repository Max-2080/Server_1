package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

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

		validate := validator.New()
		validate.RegisterValidation("no_forbidden_words", noForbiddenWords)

		ot := Feedback{}
		var response Response
		c.BodyParser(&ot)
		err := validate.Struct(ot)

		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fmt.Println("Validation failed:", validationErrors)
			response.Message = append(response.Message, validationErrors.Error())
			c.SendStatus(422)
			return c.JSON(response)
		}
		message = append(message, ot.Message)

		return c.SendString("Feedback received. Thank you, " + ot.Name + " .")
	})

	app.Listen(":3000")
}
