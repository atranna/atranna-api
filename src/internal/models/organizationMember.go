package models

type OrganizationMember struct {
	OrganizationID int `json:"organization_id" binding:"required"`
	UserID         int `json:"user_id" binding:"required"`
	Role           string `json:"role" binding:"required"`
}