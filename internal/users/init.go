package users

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
)

var OPENSTACK_DRIVER_URL string
var SECRET string

var Log *log.Logger

var APP *pocketbase.PocketBase

func Init(app *pocketbase.PocketBase, secret string) {
	Log = log.New(os.Stdout, "[user]", log.Ldate|log.Ltime|log.Llongfile)

	APP = app
	SECRET = secret

	OPENSTACK_DRIVER_URL = os.Getenv("OPENSTACK_DRIVER_URL")

	{
		req, err := http.NewRequest("GET", OPENSTACK_DRIVER_URL+"/ping", nil)
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

	app.OnUserBeforeCreateRequest().Add(UserBeforeCreate)
	app.OnUserBeforeUpdateRequest().Add(UserBeforeUpdate)
	app.OnUserBeforeDeleteRequest().Add(UserBeforeDelete)
}
