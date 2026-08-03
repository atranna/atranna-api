package postgres

import (
	"database/sql"

	"github.com/atranna/atranna-api/src/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Add(user models.User) (models.User, error) {
	err := r.db.QueryRow(
		`INSERT INTO users (email, username, password_hash, display_name) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Email, user.Username, user.PasswordHash, user.DisplayName,
	).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id int) (models.User, bool) {
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE id = $1`, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
	if err != nil {
		return models.User{}, false
	}
	return user, true
}

func (r *UserRepository) GetByUsername(username string) (models.User, bool) {
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE username = $1`, username)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
	if err != nil {
		return models.User{}, false
	}
	return user, true
}

func (r *UserRepository) GetAll() []models.User {
	rows, err := r.db.Query(`SELECT id, email, username, password_hash, display_name FROM users`)
	if err != nil {
		return []models.User{}
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return []models.User{}
	}
	return users
}

func (r *UserRepository) Delete(id int) (models.User, bool) {
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE id = $1`, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
	if err != nil {
		return models.User{}, false
	}

	_, err = r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return models.User{}, false
	}

	return user, true
}