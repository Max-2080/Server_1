package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func main1() {
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

	app.Post("/create_user", func(c *fiber.Ctx) error {

		validate := validator.New()
		ot := UserCreate{}
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

		return c.JSON(ot)
	})

	var Products = []Product{}

	app.Get("/product/:product_id", func(c *fiber.Ctx) error {
		id1 := c.Params("product_id")
		id, err := strconv.Atoi(id1)
		if err != nil {
			return c.SendStatus(400)
		}

		for i := 0; i < len(Products); i++ {
			if Products[i].Product_id == id {
				return c.JSON(Products[i])
			}
		}
		fmt.Println(Products)
		return c.SendString("Товар не найден. Для начала добавьте его")
	})

	app.Post("/prod/addproduct", func(c *fiber.Ctx) error {
		product := Product{}
		c.BodyParser(&product)
		Products = append(Products, product)
		return c.SendStatus(200)
	})

	app.Get("products/search", func(c *fiber.Ctx) error {
		var Otwet = []Product{}
		keyword := c.Query("keyword", "ERROR") // Если id не указан, используется значение "defaultID"
		if keyword == "ERROR" {
			return c.SendStatus(400)
		}
		category := c.Query("category", "null")
		limit := c.Query("limit", "1000000")
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return c.SendStatus(401)
		}
		for i := 0; i < len(Products); i++ {
			if (category == Products[i].Category || strings.Contains(Products[i].Name, keyword)) && len(Otwet) < lim {
				Otwet = append(Otwet, Products[i])
			}
		}
		return c.JSON(Otwet)
	})

	app.Listen(":3000")
}
