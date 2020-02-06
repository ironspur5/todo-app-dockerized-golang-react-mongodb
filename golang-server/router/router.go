package router

import (
	//"../middleware"
	"github.com/gorilla/mux"
	"github.com/your_github_username/your_go-todo-server_module/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")                // GET method to get all task from the DB. In Methods the first parameter is Method in this case, it is GET and second OPTIONS, this is to tackle cors .
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")               // POST method to create a task in the DB.
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")         // PUT method to update the task’s status to true in the DB. The task’s id is passed as params in the URL.
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")         // PUT method to update the task’s status to false in the DB. The task’s id is passed as params in the URL.
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")  // DELETE method to delete the task from the DB. The task’s id is passed as params in the URL.
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS") // DELETE method to delete all the tasks from the DB.
	return router                                                                                  // Return the router instance. This router will be served in the main.go
}
