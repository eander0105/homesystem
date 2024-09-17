package main

import (
    "log"
    "backend/internal/server"
    "backend/internal/database"
    "backend/internal/models"
)

func main() {
    srv := server.NewServer() // Initialize the server

    dsn := "user=user password=password dbname=mydb host=postgres port=5432 sslmode=disable"
    err := database.InitDB(dsn)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Auto-migrate the database
    err = database.DB.AutoMigrate(&models.User{}, &models.UserSession{})
    if err != nil {
        log.Fatalf("Failed to migrate the database: %v", err)
    }

    log.Println("Starting server")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
