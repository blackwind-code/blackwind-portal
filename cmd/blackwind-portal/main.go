package main

import (
	"log"
	"os"

	"github.com/blackwind-code/blackwind-portal/internal/users"
	"github.com/blackwind-code/blackwind-portal/internal/vpn"
	"github.com/pocketbase/pocketbase"
)

func getApp() *pocketbase.PocketBase {
	app := pocketbase.New()

	secret := os.Getenv("SECRET")

	users.Init(app, secret)
	vpn.Init(app, secret)

	return app
}

func main() {
	app := getApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
