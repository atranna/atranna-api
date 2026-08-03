package models

type OrganizationMember struct {
	OrganizationID int `json:"organization_id"`
	UserID         int `json:"user_id"`
}