package users

import (
	"github.com/pocketbase/pocketbase/core"
)

func UserBeforeCreate(e *core.UserCreateEvent) error {
	return nil
}

func UserBeforeUpdate(e *core.UserUpdateEvent) error {
	return nil
}

func UserBeforeDelete(e *core.UserDeleteEvent) error {
	return nil
}
