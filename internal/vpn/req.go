package vpn

type Device_Create struct {
	ZTAddress string `json:"zt_address"`
}

type Device_Update struct {
	ZTAddress  string `json:"zt_address"`
	DeviceType int    `json:"device_type"`
}

type Device_Delete struct {
	ZTAddress string `json:"zt_address"`
}
