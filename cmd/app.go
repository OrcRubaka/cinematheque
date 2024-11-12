package main

import (
	"fmt"
	"net/http"

	"cinematheque/controller"
	"cinematheque/internal/postgres"
	"cinematheque/internal/repository"
	"cinematheque/service"
)

func Run() error {
	// Подключение к базе данных
	db, err := postgres.Connect()
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	movieRepo := repository.NewMovieRepository(db)
	actorRepo := repository.NewActorRepository(db)

	// Инициализация сервисов
	movieService := service.NewMovieService(movieRepo)
	actorService := service.NewActorService(actorRepo)

	// Инициализация контроллеров
	movieController := controller.NewMovieController(movieService)
	actorController := controller.NewActorController(actorService)

	// Маршрутизация
	http.HandleFunc("/movies/create", movieController.Create)
	http.HandleFunc("/actors/create", actorController.Create)

	// Запуск HTTP-сервера
	fmt.Println("Запуск сервера на http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Ошибка при запуске приложения: %v\n", err)
	}
}
