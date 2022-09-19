package vpn

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
)

var VPN_DRIVER_ADDR string
var VPN_DRIVER_SECRET string
var Log *log.Logger

func Init(app *pocketbase.PocketBase) {

	VPN_DRIVER_ADDR = os.Getenv("VPN_DRIVER_ADDR")
	VPN_DRIVER_SECRET = os.Getenv("VPN_DRIVER_SECRET")

	/*{
		req, err := http.NewRequest("GET", VPN_DRIVER_ADDR, nil)
		if err != nil {
			Log.Fatalf("Cannot connect to driver: %v\n", err)
		}
		req.Header.Add("X-Auth-Token", VPN_DRIVER_SECRET)

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

	app.OnModelBeforeCreate().Add(DeviceBeforeCreate)
	app.OnModelBeforeUpdate().Add(DeviceBeforeUpdate)
	app.OnModelBeforeDelete().Add(DeviceBeforeDelete)
}
