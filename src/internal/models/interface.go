package models

type Interface struct {
	ID          int    `json:"id"`
	DeviceID   int    `json:"device_id" binding:"required"`
	IPAddress  string `json:"ip_address"`
	MACAddress string `json:"mac_address"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Speed       int    `json:"speed"`
}
