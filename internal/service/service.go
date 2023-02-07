package service

import (
	"context"

	"github.com/shivamk2406/Practice/internal/service/user"
)

type Registry struct {
	UserSvc user.API
}

type API interface {
	CreateUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error)
	GetUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error)
}

func (r Registry) CreateUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error) {
	return r.UserSvc.CreateUserSubScription(ctx, m)
}

func (r Registry) GetUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error) {
	return r.UserSvc.GetUserSubScription(ctx, m)
}
