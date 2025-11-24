package main

import (
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := LoadConfig()

	// DB connection & migration via Store (kept only for DB setup)
	store := NewStore(cfg)
	defer store.DB.Close()

	// Initialize logger
	initLogger()
	defer Logger.Sync()

	// Create Repository
	repo := NewMySQLStudentRepository(store.DB)

	// Create Service layer
	service := NewStudentService(repo)

	// Create Handlers
	handlers := NewStudentHandlers(service)

	// Register endpoints / routes
	http.HandleFunc("/students", LoggingMiddleware(handlers.handleStudents))
	http.HandleFunc("/students/", LoggingMiddleware(handlers.handleStudentByID))

	// Start server
	address := ":8080"
	log.Println("Server running at http://localhost" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(address, nil))
}
