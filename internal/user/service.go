package user

import (
	"github.com/tiagocesar/awesomeProject/internal/repo"
)

type userStorer interface {
	AddUser(u repo.User) string
}

type Service struct {
	userStorer userStorer
}

func NewService(storer userStorer) *Service {
	return &Service{
		storer,
	}
}

func (s *Service) AddUser(name, email string) string {
	return s.userStorer.AddUser(repo.User{
		Name:  name,
		Email: email,
	})
}
