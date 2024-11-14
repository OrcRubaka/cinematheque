package controller

import (
	"cinematheque/service"
	"encoding/json"
	"net/http"
	"strconv"
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

func (c *ActorController) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	actor, err := c.service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if actor == nil {
		http.Error(w, "Actor not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actor)
}

func (c *ActorController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	birthdate, err := time.Parse("2006-01-02", r.FormValue("birthdate"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, name, birthdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ActorController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
