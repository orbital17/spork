package spork

import (
	"spork/config"
	"spork/postgres"
	"spork/users"
)

type App struct {
	UserStore   *users.Store
	UserService *users.Service
}

func NewApp() *App {
	config := config.Default()
	db := postgres.Init(config)
	store := users.NewStore(db)
	service := users.NewService(store)
	app := App{store, service}
	return &app
}
