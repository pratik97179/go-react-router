package main

import (
	"go-react-router/internal/app"
	"log"
)

func main() {
	application, err := app.New()

	if err != nil {
		log.Fatal("failed to initialise application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatal("failed to start the server: %v", err)
	}
}
