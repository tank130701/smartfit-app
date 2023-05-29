package handler

import (
	//"my-app/internal/models"
	// "fmt"
	"my-app/internal/models"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
)
func (h *Handler) createUserData(c *gin.Context) {
	userId, err := getUserId(c)
	// fmt.Println("User ID in handler: ", userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.UserData
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// fmt.Println("User UserData in handler: ", input)
	id, err := h.services.UsersData.CreateUserData(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// fmt.Println("Hello from handler!")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}


func (h *Handler) getUserData(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	// 	return
	// }

	userData, err := h.services.UsersData.GetUserData(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userData)
}