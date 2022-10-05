package main

import (
	"log"
	"os"

	"github.com/blackwind-code/blackwind-portal/internal/users"
	"github.com/blackwind-code/blackwind-portal/internal/vpn"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

var SECRET string

func getApp() *pocketbase.PocketBase {
	app := pocketbase.New()

	SECRET = os.Getenv("SECRET")

	users.Init(app, SECRET)
	vpn.Init(app, SECRET)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("pb_public"), false))

		return nil
	})

	return app
}

func main() {
	app := getApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
