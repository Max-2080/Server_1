package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Otw struct {
	Result string
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

	app.Post("/feedback ", func(c *fiber.Ctx) error {
		ot := Feedback{}
		c.BodyParser(&ot)
		message = append(message, ot.Message)
		var response Response
		response.Message = "Feedback received. Thank you, " + ot.Name + " ."
		return c.JSON(response)
	})

	app.Listen(":3000")
}
