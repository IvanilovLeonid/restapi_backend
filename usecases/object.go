package usecases

import "hw1/models"

type Object interface {
	Get(key string) (*models.Task, error)
	Put(key string, value *models.Task) error
	Post(key string, value *models.Task) error
	Delete(key string) error
}
