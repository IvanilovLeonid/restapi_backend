package repository

import "hw1/models"

type UserRepository interface {
	Register(login, password string) error
	Authenticate(login, password string) (*models.User, error)
}
