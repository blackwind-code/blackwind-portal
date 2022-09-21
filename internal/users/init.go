package users

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
)

var OPENSTACK_DRIVER_ADDR string
var OPENSTACK_DRIVER_SECRET string
var Log *log.Logger

var APP *pocketbase.PocketBase

func Init(app *pocketbase.PocketBase) {
	Log = log.New(os.Stdout, "[user]", log.Ldate|log.Ltime|log.Llongfile)

	APP = app

	OPENSTACK_DRIVER_ADDR = os.Getenv("OPENSTACK_DRIVER_ADDR")
	OPENSTACK_DRIVER_SECRET = os.Getenv("OPENSTACK_DRIVER_SECRET")

	/*{
		req, err := http.NewRequest("GET", OPENSTACK_DRIVER_ADDR, nil)
		if err != nil {
			Log.Fatalf("Cannot connect to driver: %v\n", err)
		}
		req.Header.Add("X-Auth-Token", OPENSTACK_DRIVER_SECRET)

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
	}*/

	app.OnUserBeforeCreateRequest().Add(UserBeforeCreate)
	app.OnUserBeforeUpdateRequest().Add(UserBeforeUpdate)
	app.OnUserBeforeDeleteRequest().Add(UserBeforeDelete)
}
