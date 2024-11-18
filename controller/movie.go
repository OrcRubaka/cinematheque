package controller

import (
	"cinematheque/service"
	"github.com/gin-gonic/gin"
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

func (c *MovieController) Create(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	releaseDateStr := ctx.PostForm("release_date")

	releaseDate, err := time.Parse("2006-01-02", releaseDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	ratingStr := ctx.PostForm("rating")
	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating format"})
		return
	}

	err = c.service.Create(title, description, releaseDate, rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Movie created successfully"})
}

func (c *MovieController) Get(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	movie, err := c.service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if movie == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) Update(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	releaseDateStr := ctx.PostForm("release_date")

	releaseDate, err := time.Parse("2006-01-02", releaseDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	ratingStr := ctx.PostForm("rating")
	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating format"})
		return
	}

	err = c.service.Update(id, title, description, releaseDate, rating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func (c *MovieController) Delete(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "Movie deleted successfully"})
}
