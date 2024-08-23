package handlers

import (
    "log"
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "backend/internal/database"
)

// RegisterUser registers a new user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    log.Println("RegisterUserHandler")
    // Decode the request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Check for errors in the request
    if req.Name == "" || req.Email == "" || req.Password == "" {
        log.Println("All fields are required")
        http.Error(w, "All fields are required", http.StatusBadRequest)
        return
    }

    // Check if email is already registered
    var existingUser database.User
    if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        http.Error(w, "Email already registered", http.StatusBadRequest)
        return
    }
    // if database.DB.find($database.User{}, "email = ?", req.Email) {
    //     http.Error(w, "Email already registered", http.StatusBadRequest)
    //     return
    // }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Create the user
    user := database.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    // Save the user
    if err := database.DB.Create(&user).Error; err != nil {
        http.Error(w, "Failed to save user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
    // Login user
}
