package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type DB interface {
	CreateUserSubScription(ctx context.Context, m *Model) (*Model, error)
	GetUserSubScription(ctx context.Context, m *Model) (*Model, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) DB {
	return &Repository{db: db}
}

func (r *Repository) CreateUserSubScription(ctx context.Context, m *Model) (*Model, error) {
	result := r.db.WithContext(ctx).Create(&m)
	if err := result.Error; err != nil {
		return nil, err
	}
	return r.GetUserSubScription(ctx, &Model{ID: m.ID})
}

func (r *Repository) GetUserSubScription(ctx context.Context, m *Model) (*Model, error) {
	var res Model
	result := r.db.WithContext(ctx).
		Scopes(m.whereClause()...).
		Take(&res)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &res, nil
}
