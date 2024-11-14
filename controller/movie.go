package controller

import (
	"cinematheque/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type MovieController struct {
	service *service.MovieService
}

func NewMovieController(service *service.MovieService) *MovieController {
	return &MovieController{service: service}
}

func (c *MovieController) Create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	releaseDate, err := time.Parse("2006-01-02", r.FormValue("release_date"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		http.Error(w, "Invalid rating format", http.StatusBadRequest)
		return
	}

	err = c.service.Create(title, description, releaseDate, rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *MovieController) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	movie, err := c.service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if movie == nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (c *MovieController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	releaseDate, err := time.Parse("2006-01-02", r.FormValue("release_date"))
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		http.Error(w, "Invalid rating format", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, title, description, releaseDate, rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *MovieController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
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
