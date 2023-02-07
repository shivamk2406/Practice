package graph

import (
	"context"

	"github.com/shivamk2406/Practice/internal/service/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo user.DB
}

func NewResolver(repo user.DB) *Resolver {
	return &Resolver{
		repo: repo,
	}
}

func (r *Resolver) CreateUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error) {
	return r.repo.CreateUserSubScription(ctx, m)
}

func (r *Resolver) GetUserSubScription(ctx context.Context, m *user.Model) (*user.Model, error) {
	return r.repo.GetUserSubScription(ctx, m)
}
