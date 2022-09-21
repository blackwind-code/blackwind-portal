package main

import (
	"log"
	"os"

	"github.com/blackwind-code/blackwind-portal/internal/users"
	"github.com/blackwind-code/blackwind-portal/internal/vpn"
	"github.com/pocketbase/pocketbase"
)

var SECRET string

func getApp() *pocketbase.PocketBase {
	app := pocketbase.New()

	SECRET = os.Getenv("SECRET")

	users.Init(app, SECRET)
	vpn.Init(app, SECRET)

	return app
}

func main() {
	app := getApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
