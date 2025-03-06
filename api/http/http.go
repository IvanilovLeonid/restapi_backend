package http

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	httpSwagger "github.com/swaggo/http-swagger"
	"hw1/api/http/types"
	"hw1/usecases"
	"net/http"
)

type Server struct {
	service usecases.Object
}

func newServer(service usecases.Object) *Server {
	return &Server{service: service}
}

// postHandler creates a new task.
// @Summary Create a new task
// @Description Creates a task with an empty result and status "in_progress".
// @Tags Task
// @Accept json
// @Produce json
// @Param request body models.Task true "Task data"
// @Success 201 {object} types.PostHandlerResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /task [post]
func (db *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	// TODO task_body parsing and validation
	// creating task that empty
	req, err := types.CreatePostHandlerRequest(r)
	// realisation uuid
	req.TaskID = uuid.New().String()
	req.Status = "in_progress"
	req.Result = "not yet"
	err = db.service.Post(req.TaskID, req)
	types.ProcessErrors(w, err, types.PostHandlerResponse{TaskID: req.TaskID})

}

// getStatusHandler retrieves the status of a task.
// @Summary Get task status
// @Description Retrieves the status of a task by its ID.
// @Tags Task
// @Accept json
// @Produce json
// @Param task_id path string true "Task ID"
// @Success 200 {object} types.GetStatusHandlerResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /status/{task_id} [get]
func (db *Server) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	task, err := db.service.Get(req.TaskId)
	types.ProcessErrors(w, err, types.GetStatusHandlerResponse{Status: task.Status})
}

// getResultHandler retrieves the result of a task.
// @Summary Get task result
// @Description Retrieves the result of a task by its ID. If the task is still in progress, returns a conflict.
// @Tags Task
// @Accept json
// @Produce json
// @Param task_id path string true "Task ID"
// @Success 200 {object} types.GetResultHandlerResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Task not found"
// @Failure 409 {string} string "Task in progress"
// @Failure 500 {string} string "Internal Server Error"
// @Router /result/{task_id} [get]
func (db *Server) getResultHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetHandlerRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	task, err := db.service.Get(req.TaskId)

	if task.Status == "in_progress" {
		http.Error(w, "Task in progress", http.StatusConflict)
		task.Result = task.Code
		task.Status = "ready"
		return
	}

	types.ProcessErrors(w, err, types.GetResultHandlerResponse{Result: task.Result})
}

func CreateAndRunServer(service usecases.Object, addr string) error {
	server := newServer(service)

	r := chi.NewRouter()

	r.Post("/task", server.postHandler)

	r.Get("/status/{task_id}", server.getStatusHandler)

	r.Get("/result/{task_id}", server.getResultHandler)

	r.Get("/swagger/*", httpSwagger.Handler())

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
