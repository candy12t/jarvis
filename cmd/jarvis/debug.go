//go:build debug

package main

import (
	"log"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/handler"
)

func main() {
	if err := config.Setup(); err != nil {
		log.Fatal(err)
	}

	if err := handler.Apply(); err != nil {
		log.Fatal(err)
	}
}
