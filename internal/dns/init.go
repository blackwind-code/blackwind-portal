package dns

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
)

var CLOUDFLARE_X_AUTH_EMAIL string
var CLOUDFLARE_GLOBAL_KEY string
var DOMAIN_ZONE_IDENTIFIER string

var Log *log.Logger

var APP *pocketbase.PocketBase

func Init(app *pocketbase.PocketBase, secret string) {
	Log = log.New(os.Stdout, "[dns]", log.Ldate|log.Ltime|log.Llongfile)

	CLOUDFLARE_GLOBAL_KEY = os.Getenv("CLOUDFLARE_GLOBAL_KEY")
	Log.Printf("CLOUDFLARE_GLOBAL_KEY: %v\n", CLOUDFLARE_GLOBAL_KEY)
	CLOUDFLARE_X_AUTH_EMAIL = os.Getenv("CLOUDFLARE_X_AUTH_EMAIL")
	Log.Printf("CLOUDFLARE_X_AUTH_EMAIL: %v\n", CLOUDFLARE_X_AUTH_EMAIL)
	DOMAIN_ZONE_IDENTIFIER = os.Getenv("DOMAIN_ZONE_IDENTIFIER")
	Log.Printf("DOMAIN_ZONE_IDENTIFIER %v\n", DOMAIN_ZONE_IDENTIFIER)

	app.OnModelBeforeCreate().Add(DNSBeforeCreate)
	app.OnModelBeforeUpdate().Add(DNSBeforeUpdate)
	app.OnModelBeforeDelete().Add(DNSBeforeDelete)
}
