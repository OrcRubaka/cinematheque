package controller

import (
	"cinematheque/service"
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
