package main

type User struct {
	Id   int
	Name string
}

type Feedback struct {
	Name    string `validate:"required,min=2,max=50"`
	Message string `validate:"required,min=10,max=500,no_forbidden_words"`
}

type Response struct {
	Message []string
}

type Otw struct {
	Result string
}

type UserCreate struct {
	Name          string `validate:"required,min=2,max=50"`
	Email         string `validate:"required,email"`
	Age           int    `validate:"gt=0"`
	Is_subscribed bool   `validate:"omitempty"`
}

type Product struct {
	Product_id int
	Name       string
	Category   string
	Price      float64
}

type Login struct {
	Username string
	Password string
}
