package routes

import (
	controllers "FPPBKKGo/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupMovieRoutes(router *gin.Engine, movieController *controllers.MovieController) {
    movieGroup := router.Group("/movies")
    {
        movieGroup.GET("/:id", movieController.GetMovie)
        movieGroup.GET("/", movieController.GetAllMovies)
    }
}