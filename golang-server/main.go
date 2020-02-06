package main

import (
	"fmt"
	"log"
	"net/http"

	//"./router"
	"github.com/your_github_username/your_go-todo-server_module/router"
)

func main() {
	r := router.Router() // import router package
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r)) // http package serves/hosts routes on port 8080 and log package tracks logs
}
