package dns

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func DNSBeforeCreate(e *core.ModelEvent) error {
	if e.Model.TableName() == "dns" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "dns" {
			url := "https://api.cloudflare.com/client/v4/zones/" + DOMAIN_ZONE_IDENTIFIER + "/dns_records"

			var p DNS_Record_Create

			p.Comment = record.GetStringDataValue("email") + ";" + record.GetStringDataValue("description")
			p.Type = "A"
			p.Content = record.GetStringDataValue("ip")
			p.Name = record.GetStringDataValue("sub_domain") + "." + record.GetStringDataValue("primary_domain")
			p.Priority = 10
			p.Proxied = false
			p.Tags = []string{}
			p.TTL = 300

			reqPayload, err := json.Marshal(p)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqPayload))

			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("X-Auth-Email", CLOUDFLARE_X_AUTH_EMAIL)
			req.Header.Add("X-Auth-Key", CLOUDFLARE_GLOBAL_KEY)

			res, err := http.DefaultClient.Do(req)
			Log.Println("Sent req to Cloudflare DNS")

			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}
			Log.Println("Got res from Cloudflare DNS")

			defer res.Body.Close()
			if res.StatusCode != 200 {
				body, _ := ioutil.ReadAll(res.Body)
				Log.Printf("Error: %v\n", string(body))
				return hook.StopPropagation
			}

			var jsonBody Response_DNS_Record_Create
			body, _ := ioutil.ReadAll(res.Body)
			if err := json.Unmarshal(body, &jsonBody); err != nil {
				Log.Printf("Error: %v\n", err)
			}
			Log.Printf("CF Res: %v\n", jsonBody)
			record.SetDataValue("cloudflare_record_id", jsonBody.Result.Id)
			return nil
		}
		return hook.StopPropagation
	}
	return nil
}

func DNSBeforeUpdate(e *core.ModelEvent) error {
	if e.Model.TableName() == "dns" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "dns" {

			var p DNS_Record_Update

			p.Comment = record.GetStringDataValue("email") + ";" + record.GetStringDataValue("description")
			p.Content = record.GetStringDataValue("ip")

			url := "https://api.cloudflare.com/client/v4/zones/" + DOMAIN_ZONE_IDENTIFIER + "/dns_records/" + record.GetStringDataValue("cloudflare_record_id")

			reqPayload, err := json.Marshal(p)
			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqPayload))

			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("X-Auth-Email", CLOUDFLARE_X_AUTH_EMAIL)
			req.Header.Add("X-Auth-Key", CLOUDFLARE_GLOBAL_KEY)

			res, err := http.DefaultClient.Do(req)

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

func DNSBeforeDelete(e *core.ModelEvent) error {
	if e.Model.TableName() == "dns" {
		if record, ok := e.Model.(*models.Record); ok && record.Collection().Name == "dns" {

			url := "https://api.cloudflare.com/client/v4/zones/" + DOMAIN_ZONE_IDENTIFIER + "/dns_records/" + record.GetStringDataValue("cloudflare_record_id")

			req, err := http.NewRequest("DELETE", url, nil)

			if err != nil {
				Log.Printf("Error: %v\n", err.Error())
				return hook.StopPropagation
			}

			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("X-Auth-Email", CLOUDFLARE_X_AUTH_EMAIL)
			req.Header.Add("X-Auth-Key", CLOUDFLARE_GLOBAL_KEY)

			res, err := http.DefaultClient.Do(req)

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
