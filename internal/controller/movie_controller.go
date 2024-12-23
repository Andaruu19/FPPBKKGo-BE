package controllers

import (
	"net/http"
	"strconv"

	"FPPBKKGo/internal/usecases"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	MovieUsecase *usecases.MovieUsecase
}

// GetMovie handles fetching a movie by ID
func (mc *MovieController) GetMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	movie, err := mc.MovieUsecase.GetMovie(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// GetAllMovies fetches all movies
func (mc *MovieController) GetAllMovies(c *gin.Context) {
	movies, err := mc.MovieUsecase.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// Add to movie_controller.go
func (mc *MovieController) GetMovieBySlug(c *gin.Context) {
	slug := c.Param("slug")
	movie, err := mc.MovieUsecase.GetMovieBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) GetMovieByTitle(c *gin.Context) {
	title := c.Param("title")
	movie, err := mc.MovieUsecase.GetMovieByTitle(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) GetMovieByGenre(c *gin.Context) {
	genre := c.Param("slug")
	movie, err := mc.MovieUsecase.GetMovieByGenre(genre)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": genre, "movies": movie})
}

func (mc *MovieController) GetMovieByActor(c *gin.Context) {
	actor, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	movie, err := mc.MovieUsecase.GetMovieByActor(uint(actor))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": actor, "movies": movie})
}
