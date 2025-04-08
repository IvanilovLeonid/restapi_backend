package service

import (
	"hw1/models"
	"hw1/repository"
)

type User struct {
	usr repository.UserRepository
}

func NewUser(usr repository.UserRepository) *User {
	return &User{
		usr: usr,
	}
}

func (db *User) Register(login, password string) error { return db.usr.Register(login, password) }

func (db *User) Authenticate(login, password string) (*models.User, error) {
	return db.usr.Authenticate(login, password)
}
