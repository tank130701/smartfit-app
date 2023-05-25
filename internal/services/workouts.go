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

func (s *WorkoutsService) Get(user models.User) ([]models.Workout, error) {
	//TODO implement me
	panic("implement me")
	
}
