package controller

import (
	"cinematheque/service"
	"net/http"
	"time"
)

type ActorController struct {
	service *service.ActorService
}

func NewActorController(service *service.ActorService) *ActorController {
	return &ActorController{service: service}
}

func (c *ActorController) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	gender := r.FormValue("gender")
	birthdate, err := time.Parse("2006-01-02", r.FormValue("birthdate"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	err = c.service.Create(name, gender, birthdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
