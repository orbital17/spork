//+build wireinject

package containers

import (
	"spork/config"
	grpc "spork/grpc"
	"spork/postgres"
	"spork/rest"
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
		grpc.NewFilesServer,
		grpc.NewRunner,
		rest.NewServer,
		NewApp,
	)
	return &App{}
}
