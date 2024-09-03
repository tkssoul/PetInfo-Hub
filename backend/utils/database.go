package utils

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)



func InitDB() *gorm.DB {
    var err error
    var DB *gorm.DB
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    } 
    return DB   
}