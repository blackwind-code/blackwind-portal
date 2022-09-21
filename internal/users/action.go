package users

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func UserBeforeCreate(e *core.UserCreateEvent) error {
	// call driver to create user
	return nil
}

func UserBeforeUpdate(e *core.UserUpdateEvent) error {
	// call driver to update user
	oldEmail, _ := APP.Dao().FindUserById(e.User.Id)
	fmt.Println("Old Email: " + oldEmail.Email)
	fmt.Println("New Email: " + e.User.Email)
	return nil
}

func UserBeforeDelete(e *core.UserDeleteEvent) error {
	// call driver to delete user
	return nil
}
