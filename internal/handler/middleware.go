package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userCtx             = "userId"
)

func (h *Handler) authMiddleware(ctx *gin.Context) {
	sessionToken, err := ctx.Cookie("session_cookie")
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	session, err := h.services.Auth.GetSession(sessionToken)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	err = session.IsExpired()
	if err != nil {
		ctx.SetCookie("session_cookie", "", 0, "/", "localhost", false, true)
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		err = h.services.Auth.DeleteSession(session.ID)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		return
	}

	user, err := h.services.Auth.GetUser(session.UserID)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("user", user)
	ctx.Set("userId", user.ID)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
