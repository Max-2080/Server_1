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
