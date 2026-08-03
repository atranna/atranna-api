package models

type User struct {
	ID		     int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"-" binding:"required"`
	DisplayName  string `json:"display_name"`
}