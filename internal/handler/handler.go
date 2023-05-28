package handler

import (
	"my-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := router.Group("/api", h.authMiddleware)
	{
		public := api.Group("/public")
		{
			user := public.Group("/user")
			{
				user.GET("/", h.getUser)
				user.DELETE("/", h.deleteUser)
				user.PUT("/", h.updateUser)
			}
			workouts := public.Group("/workouts")
			{
				workouts.GET("/user", h.getUserArticles)
				workouts.POST("/new", h.saveArticle)
			}
		}
	}
	// h.initAuthRoutes(router)

	return router
}
