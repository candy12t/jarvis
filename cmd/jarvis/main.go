package main

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/handler"
)

func main() {
	setup()
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda") {
		lambda.Start(Handler)
	} else {
		Handler()
	}
}

func setup() {
	if err := config.NewTwitter(); err != nil {
		log.Fatal(err)
	}
	if err := config.NewJustwatch(); err != nil {
		log.Fatal(err)
	}
}

func Handler() {
	err := handler.Apply()
	if err != nil {
		log.Fatal(err)
	}
}
