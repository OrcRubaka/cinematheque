package service

import (
	"cinematheque/internal/repository"
	"errors"
	"time"
)

type ActorStore interface {
	Create(name string, gender string, birthdate time.Time) error
	Update(id int, name string, birthdate time.Time) error
	Get(id int) (*repository.Actor, error)
	Delete(id int) error
}

type ActorService struct {
	store ActorStore
}

func NewActorService(store ActorStore) *ActorService {
	return &ActorService{store: store}
}

func (a *ActorService) Create(name string, gender string, birthdate time.Time) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return a.store.Create(name, gender, birthdate)
}

func (a *ActorService) Update(id int, name string, birthdate time.Time) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return a.store.Update(id, name, birthdate)
}

func (a *ActorService) Get(id int) (*repository.Actor, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}
	return a.store.Get(id)
}

func (a *ActorService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return a.store.Delete(id)
}
