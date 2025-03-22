package service

import (
	"goomers-api/internal/model"
	"goomers-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUsers() ([]*model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) AddUser(user *model.User) error {
	return s.repo.Create(user)
}
func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}
func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
