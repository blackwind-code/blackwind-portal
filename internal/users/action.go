package users

import (
	"github.com/pocketbase/pocketbase/core"
)

func UserBeforeCreate(e *core.UserCreateEvent) error {
	// call driver to create user
	return nil
}

func UserBeforeUpdate(e *core.UserUpdateEvent) error {
	// call driver to update user
	return nil
}

func UserBeforeDelete(e *core.UserDeleteEvent) error {
	// call driver to delete user
	return nil
}
