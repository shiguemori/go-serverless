package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func show() (*book, error) {
	bk, err := getItem("clean code")
	if err != nil {
		return nil, err
	}

	return bk, nil
}

func main() {
	lambda.Start(show)
}
