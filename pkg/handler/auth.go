package handler

import (
	"log"
	"my-app/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"time"
)

func (h *Handler) signUp(ctx *gin.Context) {
	input := &struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Sex      bool   `json:"sex"`
		Weight   int    `json:"weight"`
		Password string `json:"password"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user := &models.User{
		Username:     input.Username,
		Age:          input.Age,
		Sex:          input.Sex,
		Weight:       input.Weight,
		PasswordHash: hash,
	}

	id, err := h.repository.SaveUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}


func (h *Handler) signIn(ctx *gin.Context) {
	input := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := ctx.BindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(err)
		return
	}

	user, err := h.repository.GetUserByUsername(input.Username)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(input.Password))
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	sessionToken := uuid.NewString()

	newSession := &models.Session{
		Session:   sessionToken,
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}

	id, err := h.repository.SaveSession(newSession)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	session, err := h.repository.GetSessionByID(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("session_cookie", sessionToken, 3600*24, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"session": session,
	})

}
