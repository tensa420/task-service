package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"task-service/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	srv, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = srv.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
