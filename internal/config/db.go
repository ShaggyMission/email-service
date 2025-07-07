package config

import (
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
    dsn := "admin:Lis.12345@tcp(user-shaggy.cd5kwh1iyrrb.us-east-1.rds.amazonaws.com:3306)/user_service_db?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = db
    log.Println("MySQL database connected with GORM")
    return DB
}
