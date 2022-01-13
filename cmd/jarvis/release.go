//go:build !debug

package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/handler"
)

func main() {
	if err := config.Setup(); err != nil {
		log.Fatal(err)
	}

	lambda.Start(handler.Apply)
}
