package sqlite

import (
	"atranna-api/src/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Add(user models.User) (models.User, error) {
	res, err := r.db.Exec(
		`INSERT INTO users (email, username, password_hash, display_name) VALUES (?, ?, ?, ?)`,
		user.Email, user.Username, user.PasswordHash, user.DisplayName,
	)
	if err != nil {
		return models.User{}, err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (r *UserRepository) GetByID(id int) (models.User, bool) {
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE id = ?`, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
	if err != nil {
		return models.User{}, false
	}
	return user, true
}

func (r *UserRepository) GetByUsername(username string) (models.User, bool) {
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE username = ?`, username)
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
	row := r.db.QueryRow(`SELECT id, email, username, password_hash, display_name FROM users WHERE id = ?`, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.DisplayName)
	if err != nil {
		return models.User{}, false
	}

	_, err = r.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return models.User{}, false
	}
	return user, true
}