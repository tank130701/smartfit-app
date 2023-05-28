package services

import (
	"my-app/internal/models"
	"my-app/internal/repository"
)

type WorkoutsService struct {
	r *repository.Repositories
}

func NewWorkoutsService(r *repository.Repositories) *WorkoutsService {
	return &WorkoutsService{r: r}
}

func (s *WorkoutsService) Generate(user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *WorkoutsService) GetWorkout(id int) (models.Workout, error) {
	workout, err := s.r.GetWorkoutByID(id)
	if err != nil {
		return models.Workout{} ,err
	}
	return workout, nil
}
