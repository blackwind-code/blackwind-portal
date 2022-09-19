package vpn

import (
	"github.com/pocketbase/pocketbase/core"
)

func DeviceBeforeCreate(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		// call driver to create device
	}

	return nil
}

func DeviceBeforeUpdate(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		// call driver to update device
	}

	return nil
}

func DeviceBeforeDelete(e *core.ModelEvent) error {
	if e.Model.TableName() == "vpn" {
		// call driver to delete device
	}

	return nil
}
