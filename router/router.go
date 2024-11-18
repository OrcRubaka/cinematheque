package router

import (
	"cinematheque/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты приложения
func SetupRouter(
	movieController *controller.MovieController,
	actorController *controller.ActorController,
) *gin.Engine {
	router := gin.Default()

	// Роуты для фильмов
	movies := router.Group("/movies")
	{
		movies.POST("/create", movieController.Create)
		movies.GET("/get/:id", movieController.Get)
		movies.PUT("/update/:id", movieController.Update)
		movies.DELETE("/delete/:id", movieController.Delete)
	}

	// Роуты для актеров
	actors := router.Group("/actors")
	{
		actors.POST("/create", actorController.Create)
		actors.GET("/get/:id", actorController.Get)
		actors.PUT("/update/:id", actorController.Update)
		actors.DELETE("/delete/:id", actorController.Delete)
	}

	return router
}
