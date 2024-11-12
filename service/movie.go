package service

import (
	"cinematheque/internal/repository"
	"errors"
	"time"
)

type MovieStore interface {
	Create(title, description string, releaseDate time.Time, rating float64) error
	Update(id int, title, description string, releaseDate time.Time, rating float64) error
	Get(id int) (*repository.Movie, error)
	Delete(id int) error
}

type MovieService struct {
	store MovieStore
}

func NewMovieService(store MovieStore) *MovieService {
	return &MovieService{store}
}

func (m *MovieService) Create(title, description string, releaseDate time.Time, rating float64) error {
	if title == "" || description == "" {
		return errors.New("title and description can't be empty")
	}
	return m.store.Create(title, description, releaseDate, rating)
}

// Обновление существующего фильма
func (m *MovieService) Update(id int, title, description string, releaseDate time.Time, rating float64) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return m.store.Update(id, title, description, releaseDate, rating)
}

// Получение фильма по ID
func (m *MovieService) Get(id int) (*repository.Movie, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}
	return m.store.Get(id)
}

// Удаление фильма по ID
func (m *MovieService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return m.store.Delete(id)
}
