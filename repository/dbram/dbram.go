package dbram

import (
	"hw1/models"
	"hw1/repository"
)

type Object struct {
	data map[string]*models.Task
}

func NewObject() *Object {
	return &Object{
		data: make(map[string]*models.Task),
	}
}

func (db *Object) Get(taskID string) (*models.Task, error) {
	if task, exists := db.data[taskID]; exists {
		return task, nil
	}
	return nil, repository.NotFound
}

func (db *Object) Put(taskID string, task *models.Task) error {
	db.data[taskID] = task
	return nil
}

func (db *Object) Post(taskID string, task *models.Task) error {
	if _, exists := db.data[taskID]; exists {
		return repository.Exist
	}
	db.data[taskID] = task
	return nil
}

func (db *Object) Delete(taskID string) error {
	if _, exists := db.data[taskID]; !exists {
		return repository.NotFound
	}
	delete(db.data, taskID)
	return nil
}
