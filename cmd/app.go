package main

import (
	"cinematheque/internal/postgres"
	"cinematheque/internal/repository"
	"fmt"
	"log"
)

func Run() error {
	db, err := postgres.Connect()
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	movieRepo := repository.NewMovie(db)
	actorRepo := repository.NewActor(db)

	fmt.Println("Репозитории инициализированы:", movieRepo, actorRepo)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Ошибка запуска приложения: %v", err)
	}
}
