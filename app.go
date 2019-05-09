package spork

import "spork/postgres"
import "spork/users"

type App struct {
	UserStore   users.Store
	UserService users.Service
}

func NewApp() *App {
	db := postgres.InitDB()
	store := users.NewStore(db)
	service := users.NewService(store)
	app := App{store, service}
	return &app
}
