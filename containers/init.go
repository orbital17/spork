//+build wireinject

package containers

import (
	"spork/config"
	grpc "spork/grpc"
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
		grpc.NewUsersServer,
		grpc.NewRunner,
		NewApp,
	)
	return &App{}
}
