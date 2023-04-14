package handler

import (
	"net/http"
	"my-app/pkg/repository"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		//auth.POST("/sign-in", h.signIn)
	}
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	})
	return router
}
