package vpn

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
)

var VPN_DRIVER_URL string
var SECRET string

var Log *log.Logger

var APP *pocketbase.PocketBase

func Init(app *pocketbase.PocketBase, secret string) {
	Log = log.New(os.Stdout, "[vpn]", log.Ldate|log.Ltime|log.Llongfile)

	VPN_DRIVER_URL = os.Getenv("VPN_DRIVER_URL")
	SECRET = secret

	APP = app

	{
		req, err := http.NewRequest("GET", VPN_DRIVER_URL+"/ping", nil)
		if err != nil {
			Log.Fatalf("Cannot connect to driver: %v\n", err)
		}
		req.Header.Add("X-Auth-Token", SECRET)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			Log.Fatalf("Cannot connect to driver: %v\n", err)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Log.Fatalf("Driver ping-pong failed: %v\n", err)
		}
		if string(data) != "pong" {
			Log.Fatalf("Driver ping-pong failed: %s\n", string(data))
		}
	}

	app.OnModelBeforeCreate().Add(DeviceBeforeCreate)
	app.OnModelBeforeUpdate().Add(DeviceBeforeUpdate)
	app.OnModelBeforeDelete().Add(DeviceBeforeDelete)
}
