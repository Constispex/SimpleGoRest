package repository

import (
	"errors"
	"prosting/backend-gin/internal/models"
)

// UserRepository simuliert einen Datenbankzugriff (hier in‑Memory).
type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{},
	}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *UserRepository) Save(user *models.User) (*models.User, error) {
	// Simpler Mechanismus, um eine ID zu vergeben
	user.ID = len(r.users) + 1
	r.users = append(r.users, *user)
	return user, nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
