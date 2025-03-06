package models

type Task struct {
	TaskID     string `json:"task_id"`
	Code       string `json:"code"`
	Compilator string `json:"compilator"`
	Status     string `json:"status"` // "in_progress" | "ready"
	Result     string `json:"result"`
}
