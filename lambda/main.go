package main

import (
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func main() {
	myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}
