package server

import (
    "net/http"
    "log"
    "github.com/gorilla/mux" // Example router package
)

// NewServer initializes and returns a new HTTP server
func NewServer() *http.Server {
    r := mux.NewRouter() // Create a new router

    // Add middleware
    r.Use(Middleware)

    SetupRoutes(r)

    // Create a new HTTP server
    srv := &http.Server{
        Addr: ":3011",
        Handler: r,
    }

    return srv
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

// Capture the status code in WriteHeader.
func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK} // Default status code

        next.ServeHTTP(rw, r)
        log.Printf("%s request: %s %d", r.Method, r.URL.Path, rw.statusCode) // Log the request and status code
    })
}