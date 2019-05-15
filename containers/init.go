//+build wireinject

package containers

import (
	"spork/config"
	"spork/postgres"
	"spork/users"

	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		config.Default,
		postgres.Init,
		users.NewStore,
		users.NewService,
		NewApp,
	)
	return &App{}
}
