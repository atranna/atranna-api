package models

type User struct {
	ID		     int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
	DisplayName  string `json:"display_name"`
}