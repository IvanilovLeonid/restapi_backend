package repository

import "errors"

var (
	NotFound = errors.New("task not found")
	Exist    = errors.New("task already exists")
)
