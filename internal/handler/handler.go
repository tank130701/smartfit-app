package handler

import (
	"my-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "my-app/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		user := api.Group("/user")
		{
			user.GET("/", h.getUser)
			// user.DELETE("/", h.deleteUser)
			// user.PUT("/", h.updateUser)
		}
		userData := api.Group("/userdata")
		{
			userData.POST("/", h.createUserData)
			userData.GET("/", h.getUserData)
			// user.DELETE("/", h.deleteUser)
			// user.PUT("/", h.updateUser)
		}
		
		workouts := api.Group("/workouts")
		{
			workouts.GET("/:id", h.getWorkoutById)
			workouts.GET("/", h.getWorkouts)
			workouts.POST("/new", h.insertWorkout)
		}
	// h.initAuthRoutes(router)
	return router
	}
}