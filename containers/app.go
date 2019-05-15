package containers

import (
	"spork/users"
)

type App struct {
	UserStore   *users.Store
	UserService *users.Service
}

func NewApp(store *users.Store, service *users.Service) *App {
	app := App{store, service}
	return &app
}

func (app *App) Run() {

}
