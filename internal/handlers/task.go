package handlers

import (
	"fmt"
	"net/http"
)

// TaskController is the default implementation of TaskHandler.
type TaskController struct {
	// dependencies like DB, logger, config can be added here
}

// NewTaskController constructs a new TaskController
func NewTaskController() *TaskController {
	return &TaskController{}
}

// List handles GET /tasks
func (t *TaskController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📋 List all tasks (placeholder)")
}

// Show handles GET /tasks/{id}
func (t *TaskController) Show(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("taskID") // chi v5 way to get path params
	fmt.Fprintf(w, "🔍 Show task with ID: %s\n", taskID)
}

// Create handles POST /tasks
func (t *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "➕ Create new task (placeholder)")
}
