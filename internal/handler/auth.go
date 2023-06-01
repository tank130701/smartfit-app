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
type LoginRequest struct {
    // Имя пользователя.
    Username string `json:"username" binding:"required"`
    // Пароль пользователя.
    Password string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Input request structure"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(ctx *gin.Context) {
	var input LoginRequest

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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Input request structure"
// @Success 200 {string} string "session"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
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

	ctx.SetCookie("session_cookie", session.Session, 3600*24, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"session": session,
		"sessionId": id,
	})
}
