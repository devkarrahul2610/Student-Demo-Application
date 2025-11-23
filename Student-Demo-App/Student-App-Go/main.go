package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()

	initLogger()        // Initialize the logger
	defer Logger.Sync() // Flush logs before exit
	defer store.DB.Close()

	// Optional seed data
	store.createStudent(Student{FirstName: "Alice", LastName: "Wong", Email: "alice@example.com", Age: 20})
	store.createStudent(Student{FirstName: "Bob", LastName: "Singh", Email: "bob@example.com", Age: 22})

	// Register endpoints / routes
	http.HandleFunc("/students", LoggingMiddleware(store.handleStudents))     // GET list, POST create
	http.HandleFunc("/students/", LoggingMiddleware(store.handleStudentByID)) // GET/PUT/DELETE by ID

	// Start server
	address := ":8080"
	log.Println("Server running at http://localhost" + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
