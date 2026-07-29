package models

type Interface struct {
	ID          int    `json:"id"`
	Device_ID   int    `json:"device_id"`
	IP_Address  string `json:"ip_address"`
	MAC_Address string `json:"mac_address"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Speed       int    `json:"speed"`
}
