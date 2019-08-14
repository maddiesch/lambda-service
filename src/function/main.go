package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event is the payload sent to lambda
type Event struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
}

func main() {
	lambda.Start(func(ctx context.Context, event Event) {
		log.Println(event.Greeting, event.Name)
	})
}
