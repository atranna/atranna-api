package models

type Device struct {
	ID       int    `json:"id" binding:"required"`
	Hostname string `json:"hostname" binding:"required"`
	IP       string `json:"ip" binding:"required"`
	Vendor   string `json:"vendor" binding:"required"`
	Model    string `json:"model" binding:"required"`
	Type     string `json:"type" binding:"required"`
}
