//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/shivamk2406/Practice/graphql/graph"
	"github.com/shivamk2406/Practice/internal/service/user"
	"gorm.io/gorm"
)

func initializedReg(db *gorm.DB) *graph.Resolver {
	wire.Build(
		user.NewRepo,
		graph.NewResolver,
	)
	return &graph.Resolver{}
}
