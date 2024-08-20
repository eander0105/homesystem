package main

import (
    "log"
    "backend/internal/server"
    "backend/internal/database"
)

func main() {
    srv := server.NewServer() // Initialize the server

    dsn := "user=user password=password dbname=mydb host=localhost port=5432 sslmode=disable"
    err := database.InitDB(dsn)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Auto-migrate the database
    err = database.DB.AutoMigrate(&database.User{},)
    if err != nil {
        log.Fatalf("Failed to migrate the database: %v", err)
    }

    log.Println("Starting server on :3011")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
