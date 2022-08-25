package user

import "context"

type API interface {
	CreateUserSubScription(ctx context.Context, m *Model) (*Model, error)
	GetUserSubScription(ctx context.Context, m *Model) (*Model, error)
}

type Service struct {
	repo DB
}

func NewService(repo DB) API {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUserSubScription(ctx context.Context, m *Model) (*Model, error) {
	return s.repo.CreateUserSubScription(ctx, m)

}

func (s *Service) GetUserSubScription(ctx context.Context, m *Model) (*Model, error) {
	return s.repo.GetUserSubScription(ctx, m)
}
