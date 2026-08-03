package models

type Organization struct {
	ID       int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Slug	string `json:"slug" binding:"required"`
}
