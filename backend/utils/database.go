package utils

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)



func InitDB() *gorm.DB {
    var err error
    var DB *gorm.DB
    dsn := "root:666@tcp(127.0.0.1:3306)/PetDB?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    } 
    return DB   
}