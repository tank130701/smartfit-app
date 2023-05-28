package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "my-app/internal/models"
)

// func (h *Handler) initAuthRoutes(router *gin.Engine) {
// 	auth := router.Group("/auth")
// 	{
// 		auth.POST("/sign-up", h.signUp)
// 		auth.POST("/sign-in", h.signIn)
// 	}
// }

func (h *Handler) signUp(ctx *gin.Context) {
	input := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Auth.CreateUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	input := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	// user, err := h.repository.GetUserByNickname(input.Username)
	// if err != nil {
	// 	newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	// 	return
	// }
	// h.services.Auth.SignIn(input.Username, input.Password)
	session, id, err := h.services.Auth.GenerateSession(input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"session": session,
		"sessionId": id,
	})
}
