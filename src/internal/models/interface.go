package models

type Interface struct {
	ID          int    `json:"id"`
	DeviceID   int    `json:"device_id" binding:"required"`
	IPAddress  string `json:"ip_address" binding:"required"`
	MACAddress string `json:"mac_address"`
	Name        string `json:"name" binding:"required"`
	State       string `json:"state"`
	Speed       int    `json:"speed"`
}
