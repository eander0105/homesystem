package models

import (
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

// User represents a user in the database
type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"unique"`
    Password string
}

func (user *User) HashPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(bytes)
    return nil
}

func (user *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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

// func (user *User) CreateSesssion() () UserSession {
    // session := UserSession{
    //     UserID: user.ID,
    // }
    // database.DB.Create(&session)
    // return session
// }

func (session *UserSession) beforeCreate(tx *gorm.DB) (err error) {
    // session.ID = generateID() // Generate a unique ID or incrementing id
    // session.Token = generateToken() // Generate a unique token
    // session.CreatedAt = time.Now().Unix()
    // session.ExpiresAt = time.Now().Add(time.Hour * 24).Unix() // How long should the session last?
    return
}
