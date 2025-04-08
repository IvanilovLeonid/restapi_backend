package dbusers

import (
	"golang.org/x/crypto/bcrypt"
	"hw1/models"
	"hw1/repository"
)

type UserDB struct {
	data map[string]*models.User
}

func NewUserDB() *UserDB {
	return &UserDB{
		data: make(map[string]*models.User),
	}
}

func (db *UserDB) Register(login, password string) error {
	if _, exists := db.data[login]; exists {
		return repository.ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	db.data[login] = &models.User{
		Id:       login,
		Login:    login,
		Password: string(hashedPassword),
	}

	return nil
}

func (db *UserDB) Authenticate(login, password string) (*models.User, error) {
	user, exists := db.data[login]
	if !exists {
		return nil, repository.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, repository.ErrInvalidPass
	}

	return user, nil
}
