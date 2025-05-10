package handlers

import "net/http"

// TaskHandler defines the behavior expected for any task controller.
// This allows easy mocking, swapping, and testing.
type TaskHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

// Compile-time check: ensure TaskController satisfies TaskHandler interface
var _ TaskHandler = (*TaskController)(nil)
