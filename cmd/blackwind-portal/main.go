package main

import (
	"log"

	"github.com/blackwind-code/blackwind-portal/internal/users"
	"github.com/blackwind-code/blackwind-portal/internal/vpn"
	"github.com/pocketbase/pocketbase"
)

func getApp() *pocketbase.PocketBase {
	app := pocketbase.New()

	users.Init(app)
	vpn.Init(app)

	return app
}

func main() {
	app := getApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
