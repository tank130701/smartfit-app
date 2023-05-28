package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"my-app/internal/models"
)

func (h *Handler) getUser(ctx *gin.Context) {
	user, ok := ctx.Keys["user"].(*models.User)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}