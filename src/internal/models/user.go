package models

type User struct {
	ID		     int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
	DisplayName  string `json:"display_name"`
	CreatedAt    string `json:"created_at" binding:"required"`
	UpdatedAt    string `json:"updated_at" binding:"required"`
}