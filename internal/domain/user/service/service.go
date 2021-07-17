package service

import (
	"possant-admin/internal/domain/user/entity"
	"possant-admin/internal/domain/user/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func UserServiceProvider(repository repository.UserRepository) UserService {
	return UserService{repository: repository}
}

func (s *UserService) FindAll() []entity.User {
	return s.repository.FindAll()
}

func (s *UserService) Save(user entity.User) {
	s.repository.Save(user)
}
