package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"my-app/internal/models"
)

// @Summary Get User
// @Security BaseAuth
// @Tags userdata
// @Description get user
// @ID get-user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user [get]
func (h *Handler) getUser(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}