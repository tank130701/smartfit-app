package services

import (
	"my-app/internal/models"
	"my-app/internal/repository"
)

type UsersDataService struct {
	r *repository.Repositories
}

func NewUsersDataService(r *repository.Repositories) *UsersDataService {
	return &UsersDataService{r: r}
}

func (s *UsersDataService) GetUserData(id int) (models.UserData, error) {
	return s.r.UsersData.Get(id)
}

func (s *UsersDataService) CreateUserData(userId int, newData models.UserData) (int, error) {
	return s.r.UsersData.Create(userId, newData)
}

func (s *UsersDataService) UpdateUserData(user models.User, newData models.UserData) error {
	//TODO implement me
	panic("implement me")
}
