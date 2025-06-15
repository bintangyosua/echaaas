package auth

import "github.com/bintangyosua/echaaas/internal/user"


type Service struct {
	Repo *Repository
}

func NewService (repo *Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) RegisterUser(name, email string) error {
	user := &user.User{
		Name:  name,
		Email: email,
	}

	return s.Repo.Create(user)
}