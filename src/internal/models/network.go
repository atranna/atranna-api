package models

type Network struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	CIDR    string `json:"cidr" binding:"required"`
	Gateway string `json:"gateway" binding:"required"`
	Vlan    int    `json:"vlan"`
}
