package memory

import "atranna-api/src/internal/models"

type UserRepository struct {
	users []models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Add(user models.User) (models.User, error) {
	r.nextID++
	user.ID = r.nextID
	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepository) GetByID(id int) (models.User, bool) {
	for _, user := range r.users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

func (r *UserRepository) GetAll() []models.User {
	return r.users
}

func (r *UserRepository) Delete(id int) (models.User, bool) {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return user, true
		}
	}
	return models.User{}, false
}