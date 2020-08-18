package main

import (
	"crawlerdatacovid/internal/pkg/router"
)

func main() {
	app := router.NewApp()
	app.Run()
}
