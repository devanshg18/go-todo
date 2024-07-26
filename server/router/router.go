package router

import (
	"github.com/devanshg18/go-todo/server/middleware" // Import the middleware package
	"github.com/gorilla/mux"
)

// Router initializes and returns a new router
func Router() *mux.Router {
	router := mux.NewRouter()

	// Define routes with correct methods
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")           // Fixed path and method
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")    // Fixed path and method
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS") // Fixed path and method

	return router
}
