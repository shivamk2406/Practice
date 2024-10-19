//go:build wireinject
// +build wireinject

package app

// wire.go
import (
	"github.com/google/wire"
	userActionJson "github.com/shivamk2406/Practice/practice/datatypes/adapters/json"
	userActionCache "github.com/shivamk2406/Practice/practice/datatypes/adapters/memcache"
	userAction "github.com/shivamk2406/Practice/practice/datatypes/interfaces"
	userActionSvc "github.com/shivamk2406/Practice/practice/datatypes/service"
)

func initializedReg() userAction.API {

	wire.Build(
		userActionJson.NewUserActionJson,
		userActionCache.NewUserActionCache,
		userActionSvc.NewUserActionService,
	)
	return &userActionSvc.UserActionService{}
}
