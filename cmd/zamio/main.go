package main

import (
	"context"
	"flag"
	"log"
	"os"
	"zamio/internal/repository"
	"zamio/internal/repository/postgres"
	"zamio/internal/service"

	"github.com/joho/godotenv"
)

var (
	emailFirst  string
	emailSecond string
	sum         int
)

type zamio struct {
	repo     repository.Interface
	services *service.Services
}

func main() {
	flag.StringVar(&emailFirst, "email-first", "", "аккаунт А")
	flag.StringVar(&emailSecond, "email-second", "", "аккаунт Б")
	flag.IntVar(&sum, "sum", 0, "сумма перевода")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("loading env variables error: %v", err)
	}

	repo, err := postgres.New(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		false,
	)
	if err != nil {
		log.Fatalf("failed to establish a connection with the database: %v", err)
	}

	services := service.New(repo)

	app := &zamio{
		repo:     repo,
		services: services,
	}

	err = app.services.Payments.ProcessTransaction(context.Background(), emailFirst, emailSecond, sum)
	if err != nil {
		log.Fatal(err)
	}
}
