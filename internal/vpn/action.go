package vpn

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func DeviceBeforeCreate(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "vpn" {
			ztAddr := record.GetStringDataValue("zerotier_address")

			var p Device_Create
			p.ZTAddress = ztAddr
			reqPayload, err := json.Marshal(p)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			URL := VPN_DRIVER_URL + "/api/zerotier/device"
			req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqPayload))
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			req.Header.Set("X-Auth-Token", SECRET)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				body, _ := ioutil.ReadAll(res.Body)
				Log.Printf("Error: %v\n", string(body))
				return hook.StopPropagation
			}

			return nil
		}

		return hook.StopPropagation
	}

	return nil
}

func DeviceBeforeUpdate(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "vpn" {
			ztAddr := record.GetStringDataValue("zerotier_address")
			ztTypeStr := record.GetStringDataValue("device_type")
			ztType := 130

			if ztTypeStr == "server" {
				ztType = 120
			}

			var p Device_Update
			p.ZTAddress = ztAddr
			p.DeviceType = ztType
			reqPayload, err := json.Marshal(p)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			URL := VPN_DRIVER_URL + "/api/zerotier/device"
			req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(reqPayload))
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			req.Header.Set("X-Auth-Token", SECRET)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				body, _ := ioutil.ReadAll(res.Body)
				Log.Printf("Error: %v\n", string(body))
				return hook.StopPropagation
			}

			return nil
		}

		return hook.StopPropagation
	}

	return nil
}

func DeviceBeforeDelete(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "vpn" {
			ztAddr := record.GetStringDataValue("zerotier_address")

			var p Device_Create
			p.ZTAddress = ztAddr
			reqPayload, err := json.Marshal(p)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			URL := VPN_DRIVER_URL + "/api/zerotier/device"
			req, err := http.NewRequest("DELETE", URL, bytes.NewBuffer(reqPayload))
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			req.Header.Set("X-Auth-Token", SECRET)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				body, _ := ioutil.ReadAll(res.Body)
				Log.Printf("Error: %v\n", string(body))
				return hook.StopPropagation
			}

			return nil
		}

		return hook.StopPropagation
	}

	return nil
}
