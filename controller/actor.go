package controller

import (
	"cinematheque/service"
	"github.com/gin-gonic/gin"
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

func (c *ActorController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	gender := ctx.PostForm("gender")
	birthdateStr := ctx.PostForm("birthdate")

	birthdate, err := time.Parse("2006-01-02", birthdateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	err = c.service.Create(name, gender, birthdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Actor created successfully"})
}

func (c *ActorController) Get(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	actor, err := c.service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if actor == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	ctx.JSON(http.StatusOK, actor)
}

func (c *ActorController) Update(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	name := ctx.PostForm("name")
	birthdateStr := ctx.PostForm("birthdate")

	birthdate, err := time.Parse("2006-01-02", birthdateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	err = c.service.Update(id, name, birthdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Actor updated successfully"})
}

func (c *ActorController) Delete(ctx *gin.Context) {
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

	ctx.JSON(http.StatusNoContent, gin.H{"message": "Actor deleted successfully"})
}
