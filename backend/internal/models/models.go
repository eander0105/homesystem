package models

import (
    "gorm.io/gorm"
)

// User represents a user in the database
type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"unique"`
    Password string
}

// UserSession represents a user session in the database
type UserSession struct {
    ID        uint
    CreatedAt int64
    UserID    uint
    User      User
    Token     string `gorm:"unique"`
    ExpiresAt int64
}
