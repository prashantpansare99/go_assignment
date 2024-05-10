package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var db *sql.DB

// InitializeDB initializes the database connection.
func InitializeDB(dataSourceName string) error {
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }
    return db.Ping()
}

// CloseDB closes the database connection.
func CloseDB() error {
    return db.Close()
}