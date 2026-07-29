package models

type Device struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Vendor   string `json:"vendor"`
	Model    string `json:"model"`
	Type     string `json:"type"`
	LastSeen int64  `json:"last_seen"`
}
