//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/shivamk2406/Practice/internal/service/user"
	"gorm.io/gorm"
)

func initializedReg(db *gorm.DB) user.API {
	wire.Build(
		user.NewRepo,
		user.NewService,
	)
	return &user.Service{}
}
