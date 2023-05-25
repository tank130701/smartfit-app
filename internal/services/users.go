package services

import (
	"my-app/internal/models"
	"my-app/internal/repository"
)

type UsersService struct {
	r *repository.Repositories
}

func NewUsersService(r *repository.Repositories) *UsersService {
	return &UsersService{r: r}
}

func (s *UsersService) Get(session models.Session) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UsersService) EditData(user models.User, newData models.UserData) error {
	//TODO implement me
	panic("implement me")
}
