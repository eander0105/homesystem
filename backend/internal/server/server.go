package server

import (
    "net/http"
    "log"
    "github.com/gorilla/mux" // Example router package
    "github.com/rs/cors"
)

// NewServer initializes and returns a new HTTP server
func NewServer() *http.Server {
    r := mux.NewRouter() // Create a new router

    // Handle CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"},
        AllowCredentials: true
    })
    
    // Add middleware
    r.Use(Middleware)
    
    SetupRoutes(r)
    handler := c.Handler(r)
    
    // Create a new HTTP server
    srv := &http.Server{
        Addr: ":3011",
        Handler: handler,
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

func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("AccesswControl-Allow-Headers", "Content-Type, Authorization")
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK} // Default status code
        enableCORS(w)

        log.Printf("%s request: %s", r.Method, r.URL.Path) // Log the request
        next.ServeHTTP(rw, r)
        log.Printf("%s request: %s %d", r.Method, r.URL.Path, rw.statusCode) // Log the request and status code
    })
}