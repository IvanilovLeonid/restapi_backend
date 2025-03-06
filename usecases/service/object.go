package service

import (
	"hw1/models"
	"hw1/repository"
)

type Object struct {
	repo repository.Object
}

func NewObject(repo repository.Object) *Object {
	return &Object{
		repo: repo,
	}
}

func (db *Object) Get(taskID string) (*models.Task, error) {
	return db.repo.Get(taskID)
}

func (db *Object) Put(taskID string, task *models.Task) error {
	return db.repo.Put(taskID, task)
}

func (db *Object) Post(taskID string, task *models.Task) error {
	return db.repo.Post(taskID, task)
}

func (db *Object) Delete(taskID string) error {
	return db.repo.Delete(taskID)
}
