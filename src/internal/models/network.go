package models

type Network struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	CIDR    string `json:"cidr" binding:"required"`
	Gateway string `json:"gateway" binding:"required"`
	Vlan    int    `json:"vlan"`
}
