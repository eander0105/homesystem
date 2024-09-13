package handlers

import (
    "log"
    "net/http"
    "encoding/json"
    "backend/internal/database"
    "backend/internal/models"
    e "backend/internal/errors"
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
        e.InvalidRequestError(w)
        return
    }

    // Check if email is already registered
    // TODO: this seems to cast an error when not finding user (intended)
    var existingUser models.User
    if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        e.BadRequestError(w, "Email already registered")
        return
    }

    // Create the user
    user := models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: "",
    }

    // Hash the password
    if err := user.HashPassword(req.Password); err != nil {
        e.InternalServerError(w, err)
        return
    }

    // Save the user
    if err := database.DB.Create(&user).Error; err != nil {
        e.InternalServerError(w, err)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Decode the request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        e.InvalidRequestError(w)
        return
    }

    // Find user
    var user models.User
    if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        e.BadRequestError(w, "User not found")
        return
    }

    // Check password
    if err := user.CheckPassword(req.Password); err != nil {
        e.BadRequestError(w, "Invalid password")
        return
    }

    // Create session
    session := models.UserSession{
        UserID: user.ID,
    }

    log.Println(session)
}
