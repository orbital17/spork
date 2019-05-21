package containers

import (
	grpc "spork/grpc"
	"spork/rest"
)

type App struct {
	Grpc *grpc.Runner
	Rest *rest.Server
}

func NewApp(grpc *grpc.Runner, rest *rest.Server) *App {
	app := App{grpc, rest}
	return &app
}

func (app *App) Run() {
	// go app.Grpc.Run()
	app.Rest.Serve()
}
