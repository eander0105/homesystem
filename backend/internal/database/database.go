package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DB is the global database connection pool
var DB *gorm.DB

// InitDB initializes the database connection using GORM
func InitDB(dsn string) error {
    var err error

    // Connect to PostgreSQL using GORM
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    log.Println("Database connection established")
    return nil
}
