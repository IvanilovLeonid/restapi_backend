package repository

import "errors"

var (
	NotFound        = errors.New("task not found")
	Exist           = errors.New("task already exists")
	ErrUserExists   = errors.New("user already esxists")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPass  = errors.New("invalid password")
)
