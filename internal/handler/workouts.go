package handler

import (
	"my-app/internal/models"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
)
type getWorkoutsResponse struct {
	Data []models.Workout `json:"data"`
}

func (h *Handler) insertWorkout(c *gin.Context) {
	var input models.Workout
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Workouts.InsertWorkout(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getWorkoutById(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	workout, err := h.services.Workouts.GetWorkout(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, workout)
}

// @Summary Get Workouts
// @Security BaseAuth
// @Tags workouts
// @Description get workouts 
// @ID get-workouts
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Workout
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/workouts [get]
func (h *Handler) getWorkouts(c *gin.Context) {
	
	workouts, err := h.services.Workouts.GetWorkouts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getWorkoutsResponse{
		Data: workouts,
	})
}