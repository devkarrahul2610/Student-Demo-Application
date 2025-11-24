package main

import (
	"log"
	"net/http"
	"student/student-demo-app/internal/config"
	"student/student-demo-app/internal/database"
	"student/student-demo-app/logger"
	"student/student-demo-app/middleware"
	"student/student-demo-app/student"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// DB connection & migration via Store (kept only for DB setup)
	store := database.NewStore(cfg)
	defer store.DB.Close()

	// Initialize logger
	logger.InitLogger()
	defer logger.Logger.Sync()

	// Create Repository
	repo := student.NewMySQLStudentRepository(store.DB)

	// Create Service layer
	service := student.NewStudentService(repo)

	// Create Handlers
	handlers := student.NewStudentHandlers(service)

	// Register endpoints / routes
	http.HandleFunc("/students", middleware.LoggingMiddleware(handlers.HandleStudents))
	http.HandleFunc("/students/", middleware.LoggingMiddleware(handlers.HandleStudentByID))

	// Start server
	address := ":8080"
	log.Println("Server running at http://localhost" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(address, nil))
}
