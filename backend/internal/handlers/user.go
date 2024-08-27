package handlers

import (
    // "log"
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "backend/internal/database"
    "backend/internal/models"
    "backend/internal/errors"
)

// RegisterUser registers a new user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Decode the request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        errors.InvalidRequestError(w)
        return
    }

    // Check if email is already registered
    var existingUser models.User
    if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        errors.BadRequestError(w, "Email already registered")
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Create the user
    user := models.User{
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
