package models

type Task struct {
	TaskID     string `json:"task_id"`
	Code       string `json:"code"`
	Compilator string `json:"compilator"`
	Status     string `json:"status"` // "in_progress" | "ready"
	Result     string `json:"result"`
}

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type Session struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}
