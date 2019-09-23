package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event is the payload sent to lambda
type Event struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(func(ctx context.Context, event Event) {
		log.Println(greeting(event))
	})
}

func greeting(event Event) string {
	return fmt.Sprintf("%s %s!", os.Getenv("GREETING"), event.Name)
}
