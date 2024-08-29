package utils

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }    
}