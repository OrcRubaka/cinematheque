package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrations(db *sql.DB) {
	// Подключаемся к базе для миграций
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating migration driver: %v", err)
	}

	// Путь до папки миграций
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Migration setup error: %v", err)
	}

	// Запуск миграций
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration error: %v", err)
	}

	log.Println("Migrations applied successfully.")
}
