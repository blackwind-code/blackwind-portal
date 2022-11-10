package users

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func UserBeforeCreate(e *core.UserCreateEvent) error {
	var p OS_User_Create
	p.Email = e.User.Email
	p.PassworHash = e.User.PasswordHash
	reqPayload, err := json.Marshal(p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return hook.StopPropagation
	}

	URL := OPENSTACK_DRIVER_URL + "/api/openstack/user"
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

func UserBeforeUpdate(e *core.UserUpdateEvent) error {
	var p OS_User_Update
	oldUser, _ := APP.Dao().FindUserById(e.User.Id)
	p.OldEmail = oldUser.Email
	p.NewEmail = e.User.Email
	p.PasswordHash = e.User.PasswordHash
	reqPayload, err := json.Marshal(p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return hook.StopPropagation
	}

	URL := OPENSTACK_DRIVER_URL + "/api/openstack/user"
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

func UserBeforePasswordReset(e *core.ModelEvent) error {
	if e.Model.TableName() != "_users" {
		return nil
	}

	user := e.Model.(*models.User)
	oldUser, err := APP.Dao().FindUserById(user.Id)
	if err != nil {
		return err
	}

	// password has been changed
	if oldUser.PasswordHash != user.PasswordHash {
		var p OS_User_Update

		p.OldEmail = oldUser.Email
		p.NewEmail = user.Email
		p.PasswordHash = user.PasswordHash

		reqPayload, err := json.Marshal(p)
		if err != nil {
			Log.Printf("Error: %v\n", err.Error())
			return hook.StopPropagation
		}

		URL := OPENSTACK_DRIVER_URL + "/api/openstack/user"
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
	}

	return nil
}

func UserBeforeDelete(e *core.UserDeleteEvent) error {
	var p OS_User_Delete
	p.Email = e.User.Email
	reqPayload, err := json.Marshal(p)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return hook.StopPropagation
	}

	URL := OPENSTACK_DRIVER_URL + "/api/openstack/user"
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
