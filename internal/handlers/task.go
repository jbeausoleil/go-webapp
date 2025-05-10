package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// TaskController is the default implementation of TaskHandler.
type TaskController struct {
	// dependencies like DB, logger, config can be added here
	Tmpl *template.Template
}

// NewTaskController constructs a new TaskController
func NewTaskController(tmpl *template.Template) *TaskController {
	return &TaskController{
		Tmpl: tmpl,
	}
}

// List handles GET /tasks
func (t *TaskController) List(w http.ResponseWriter, r *http.Request) {
	data := struct {
		User string
	}{
		User: "Justin",
	}
	err := t.Tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, "render error", 500)
	}
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
