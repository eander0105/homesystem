package server

import (
    "github.com/gorilla/mux" // Example router package
    "net/http"
    "errors"
    "backend/internal/handlers" // Import your handlers
)

// SetupRoutes configures the application's routes
func SetupRoutes(r *mux.Router) {
    // Add more routes as needed
    r.HandleFunc("/hello", helloHandler).Methods("GET")
    r.HandleFunc("/error", errorHandler).Methods("GET")
    r.HandleFunc("/user/register", handlers.RegisterUserHandler).Methods("POST")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
    // Emulate an error
    err := errors.New("an error occurred")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("This will never be reached"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
