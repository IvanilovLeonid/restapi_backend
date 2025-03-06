package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"hw1/models"
	"hw1/repository"
	"net/http"
)

type GetHandlerRequest struct {
	TaskId string `json:"task_id"`
}
type GetResultHandlerResponse struct {
	Result string `json:"result"`
}
type PostHandlerResponse struct {
	TaskID string `json:"task_id"`
}

func CreatePostHandlerRequest(r *http.Request) (*models.Task, error) {
	var req models.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("Error while deceoding %v", err)
	}

	return &req, nil
}

func CreateGetHandlerRequest(r *http.Request) (*GetHandlerRequest, error) {
	taskID := chi.URLParam(r, "task_id")
	if taskID == "" {
		return nil, errors.New("missing task_id in params")
	}

	return &GetHandlerRequest{TaskId: taskID}, nil
}

type GetStatusHandlerResponse struct {
	Status string `json:"status"`
}

func ProcessErrors(w http.ResponseWriter, err error, resp any) {
	if err == repository.NotFound {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
}
