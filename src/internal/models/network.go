package models

type Network struct {
	ID       int `json:"id"`
	Name      string `json:"name"`
	CIDR      string `json:"cidr"`
	Gateway   string `json:"gateway"`
	Vlan      int    `json:"vlan"`
}
