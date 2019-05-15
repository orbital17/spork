package spork

import "spork/containers"

func Main() {
	app := containers.InitApp()
	app.Run()
}
