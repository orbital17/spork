package containers

import (
	grpc "spork/grpc"
)

type App struct {
	GrpcRunner *grpc.Runner
}

func NewApp(grpcRunner *grpc.Runner) *App {
	app := App{grpcRunner}
	return &app
}

func (app *App) Run() {
	app.GrpcRunner.Run()
}
